#ifndef HEXLIB_HEXLIB_H
#define HEXLIB_HEXLIB_H

#ifdef __cplusplus
extern "C" {
#endif

#if defined(_WIN32) || defined(_WIN64)
  #if defined(HEXLIB_EXPORTS)
    #define HEXLIB_API __declspec(dllexport)
  #else
    #define HEXLIB_API __declspec(dllimport)
  #endif
#else
  #define HEXLIB_API
#endif

// Basic functions
HEXLIB_API int hex_add(int a, int b);
HEXLIB_API const char* hex_version(void);

// MPC Kalman Filter prediction system
typedef struct {
    double value;
    double uncertainty;
} PredictionResult;

typedef struct {
    PredictionResult* results;
    int count;
} PredictionArray;

// Initialize MPC predictor with historical data
HEXLIB_API void* mpc_create_predictor(const double* historical_data, int data_count);

// Run prediction for specified steps
HEXLIB_API PredictionArray* mpc_predict(void* predictor, int prediction_steps);

// Update predictor with new measurement
HEXLIB_API void mpc_update_measurement(void* predictor, double measurement);

// Get current filtered value
HEXLIB_API PredictionResult mpc_get_current_state(void* predictor);

// Cleanup
HEXLIB_API void mpc_free_predictor(void* predictor);
HEXLIB_API void mpc_free_prediction_array(PredictionArray* array);

#ifdef __cplusplus
}
#endif

#endif /* HEXLIB_HEXLIB_H */


