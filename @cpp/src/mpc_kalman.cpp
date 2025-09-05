#include "hexlib/hexlib.h"
#include <vector>
#include <cmath>
#include <algorithm>

// 简化的3x3矩阵类
struct Matrix3x3 {
    double data[9];
    
    Matrix3x3() {
        for(int i = 0; i < 9; i++) data[i] = 0.0;
    }
    
    double& operator()(int i, int j) { return data[i*3 + j]; }
    const double& operator()(int i, int j) const { return data[i*3 + j]; }
    
    Matrix3x3 operator*(const Matrix3x3& other) const {
        Matrix3x3 result;
        for(int i = 0; i < 3; i++) {
            for(int j = 0; j < 3; j++) {
                for(int k = 0; k < 3; k++) {
                    result(i,j) += (*this)(i,k) * other(k,j);
                }
            }
        }
        return result;
    }
    
    void setIdentity() {
        for(int i = 0; i < 9; i++) data[i] = 0.0;
        data[0] = data[4] = data[8] = 1.0;
    }
    
    Matrix3x3 transpose() const {
        Matrix3x3 result;
        for(int i = 0; i < 3; i++) {
            for(int j = 0; j < 3; j++) {
                result(j,i) = (*this)(i,j);
            }
        }
        return result;
    }
    
    Matrix3x3 operator+(const Matrix3x3& other) const {
        Matrix3x3 result;
        for(int i = 0; i < 9; i++) {
            result.data[i] = data[i] + other.data[i];
        }
        return result;
    }
    
    Matrix3x3 operator-(const Matrix3x3& other) const {
        Matrix3x3 result;
        for(int i = 0; i < 9; i++) {
            result.data[i] = data[i] - other.data[i];
        }
        return result;
    }
};

// 3维向量
struct Vector3 {
    double x, y, z;
    Vector3(double x = 0, double y = 0, double z = 0) : x(x), y(y), z(z) {}
    
    Vector3 operator+(const Vector3& other) const {
        return Vector3(x + other.x, y + other.y, z + other.z);
    }
    
    Vector3 operator*(double scalar) const {
        return Vector3(x * scalar, y * scalar, z * scalar);
    }
};

// MPC卡尔曼滤波器类
class MPCKalmanFilter {
private:
    Vector3 state_posterior;         // 后验状态 [值, 速度, 加速度]
    Matrix3x3 covariance_posterior;  // 后验协方差
    Vector3 state_prior;             // 先验状态
    Matrix3x3 covariance_prior;      // 先验协方差
    Matrix3x3 transition_matrix;     // 状态转移矩阵
    Matrix3x3 process_noise_cov;     // 过程噪声协方差
    double measurement_noise_var;    // 测量噪声方差
    double time_step;
    
public:
    MPCKalmanFilter() : time_step(1.0), measurement_noise_var(1.0) {
        configureMicrogridModel();
    }
    
    void configureMicrogridModel() {
        double dt = time_step;
        double trend_decay = 0.25;
        
        // 状态转移矩阵 F
        transition_matrix.setIdentity();
        transition_matrix(0,1) = dt;
        transition_matrix(0,2) = 0.5 * dt * dt;
        transition_matrix(1,1) = 1.0 - trend_decay * dt;
        transition_matrix(1,2) = dt;
        transition_matrix(2,2) = 1.0 - 2 * trend_decay * dt;
        
        // 过程噪声协方差 Q
        process_noise_cov = Matrix3x3();
        process_noise_cov(0,0) = 0.05;
        process_noise_cov(1,1) = 0.3;
        process_noise_cov(2,2) = 0.1;
        
        // 初始协方差
        covariance_posterior.setIdentity();
        for(int i = 0; i < 3; i++) {
            covariance_posterior(i,i) = 5.0;
        }
    }
    
    void initializeFromData(const std::vector<double>& data) {
        if(data.size() < 3) return;
        
        double current_value = data.back();
        double velocity = (data[data.size()-1] - data[data.size()-2]) / time_step;
        double acceleration = 0.0;
        
        if(data.size() >= 3) {
            double vel_prev = (data[data.size()-2] - data[data.size()-3]) / time_step;
            acceleration = (velocity - vel_prev) / time_step;
        }
        
        state_posterior = Vector3(current_value, velocity, acceleration);
    }
    
    Vector3 predict() {
        // 状态预测: x_prior = F * x_posterior
        state_prior.x = transition_matrix(0,0) * state_posterior.x + 
                       transition_matrix(0,1) * state_posterior.y + 
                       transition_matrix(0,2) * state_posterior.z;
        state_prior.y = transition_matrix(1,0) * state_posterior.x + 
                       transition_matrix(1,1) * state_posterior.y + 
                       transition_matrix(1,2) * state_posterior.z;
        state_prior.z = transition_matrix(2,0) * state_posterior.x + 
                       transition_matrix(2,1) * state_posterior.y + 
                       transition_matrix(2,2) * state_posterior.z;
        
        // 协方差预测: P_prior = F * P_posterior * F^T + Q
        Matrix3x3 temp = transition_matrix * covariance_posterior;
        covariance_prior = temp * transition_matrix.transpose();
        covariance_prior = covariance_prior + process_noise_cov;
        
        // 非负约束: 确保预测值不为负数
        if (state_prior.x < 0.0) {
            state_prior.x = 0.0;
        }
        
        return state_prior;
    }
    
    void update(double measurement) {
        // 创新 (innovation): y = z - H * x_prior
        // 对于单测量值，H = [1, 0, 0]
        double innovation = measurement - state_prior.x;
        
        // 创新协方差: S = H * P_prior * H^T + R
        double innovation_cov = covariance_prior(0,0) + measurement_noise_var;
        
        // 卡尔曼增益: K = P_prior * H^T * S^(-1)
        Vector3 kalman_gain;
        kalman_gain.x = covariance_prior(0,0) / innovation_cov;
        kalman_gain.y = covariance_prior(1,0) / innovation_cov;
        kalman_gain.z = covariance_prior(2,0) / innovation_cov;
        
        // 后验状态更新: x_posterior = x_prior + K * y
        state_posterior.x = state_prior.x + kalman_gain.x * innovation;
        state_posterior.y = state_prior.y + kalman_gain.y * innovation;
        state_posterior.z = state_prior.z + kalman_gain.z * innovation;
        
        // 协方差更新: P_posterior = (I - K * H) * P_prior
        // 对于H = [1, 0, 0]，K*H是一个矩阵，第一列是K，其余为0
        Matrix3x3 I_minus_KH;
        I_minus_KH.setIdentity();
        I_minus_KH(0,0) -= kalman_gain.x;
        I_minus_KH(1,0) -= kalman_gain.y;
        I_minus_KH(2,0) -= kalman_gain.z;
        
        covariance_posterior = I_minus_KH * covariance_prior;
        
        // 非负约束: 确保更新后的状态值不为负数
        if (state_posterior.x < 0.0) {
            state_posterior.x = 0.0;
        }
    }
    
    std::vector<PredictionResult> predictFuture(int steps) {
        std::vector<PredictionResult> results;
        
        Vector3 temp_state = state_posterior;
        Matrix3x3 temp_cov = covariance_posterior;
        
        for(int i = 0; i < steps; i++) {
            // 状态预测: x = F * x
            Vector3 next_state;
            next_state.x = transition_matrix(0,0) * temp_state.x + 
                          transition_matrix(0,1) * temp_state.y + 
                          transition_matrix(0,2) * temp_state.z;
            next_state.y = transition_matrix(1,0) * temp_state.x + 
                          transition_matrix(1,1) * temp_state.y + 
                          transition_matrix(1,2) * temp_state.z;
            next_state.z = transition_matrix(2,0) * temp_state.x + 
                          transition_matrix(2,1) * temp_state.y + 
                          transition_matrix(2,2) * temp_state.z;
            
            // 协方差预测: P = F * P * F^T + Q
            Matrix3x3 temp_matrix = transition_matrix * temp_cov;
            temp_cov = temp_matrix * transition_matrix.transpose();
            temp_cov = temp_cov + process_noise_cov;
            
            temp_state = next_state;
            
            // 非负约束: 确保预测值不为负数
            if (temp_state.x < 0.0) {
                temp_state.x = 0.0;
            }
            
            PredictionResult result;
            result.value = temp_state.x;
            result.uncertainty = std::sqrt(std::max(0.0, temp_cov(0,0)));
            
            results.push_back(result);
        }
        
        return results;
    }
    
    PredictionResult getCurrentState() {
        PredictionResult result;
        result.value = state_posterior.x;
        result.uncertainty = std::sqrt(std::max(0.0, covariance_posterior(0,0)));
        return result;
    }
};

// C API实现
extern "C" {

void* mpc_create_predictor(const double* historical_data, int data_count) {
    if(!historical_data || data_count < 3) return nullptr;
    
    MPCKalmanFilter* filter = new MPCKalmanFilter();
    
    std::vector<double> data(historical_data, historical_data + data_count);
    filter->initializeFromData(data);
    
    // 在线学习阶段 - 对每个历史数据点进行预测和更新
    for(const double& measurement : data) {
        filter->predict();
        filter->update(measurement);
    }
    
    return static_cast<void*>(filter);
}

PredictionArray* mpc_predict(void* predictor, int prediction_steps) {
    if(!predictor || prediction_steps <= 0) return nullptr;
    
    MPCKalmanFilter* filter = static_cast<MPCKalmanFilter*>(predictor);
    std::vector<PredictionResult> results = filter->predictFuture(prediction_steps);
    
    PredictionArray* array = new PredictionArray();
    array->count = results.size();
    array->results = new PredictionResult[array->count];
    
    for(int i = 0; i < array->count; i++) {
        array->results[i] = results[i];
    }
    
    return array;
}

void mpc_update_measurement(void* predictor, double measurement) {
    if(!predictor) return;
    
    MPCKalmanFilter* filter = static_cast<MPCKalmanFilter*>(predictor);
    filter->predict();
    filter->update(measurement);
}

PredictionResult mpc_get_current_state(void* predictor) {
    PredictionResult result = {0.0, 0.0};
    if(!predictor) return result;
    
    MPCKalmanFilter* filter = static_cast<MPCKalmanFilter*>(predictor);
    return filter->getCurrentState();
}

void mpc_free_predictor(void* predictor) {
    if(predictor) {
        delete static_cast<MPCKalmanFilter*>(predictor);
    }
}

void mpc_free_prediction_array(PredictionArray* array) {
    if(array) {
        delete[] array->results;
        delete array;
    }
}

} // extern "C"
