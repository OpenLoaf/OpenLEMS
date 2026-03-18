package hexlib

/*
#cgo CXXFLAGS: -std=c++11
#cgo CFLAGS: -I./src/include
#cgo windows CFLAGS: -DHEXLIB_EXPORTS
#cgo windows CXXFLAGS: -DHEXLIB_EXPORTS
#cgo darwin LDFLAGS: -L./build -lhexlib -lc++
#cgo linux LDFLAGS: -L./build -lhexlib -lstdc++
#cgo windows LDFLAGS: -L./build -lhexlib
#include "hexlib.h"
*/
import "C"

import (
	"unsafe"
)

// HexAdd 执行基础加法运算
func HexAdd(a, b int) int {
	return int(C.hex_add(C.int(a), C.int(b)))
}

// HexVersion 获取库版本信息
func HexVersion() string {
	version := C.hex_version()
	return C.GoString(version)
}

// PredictionResult 预测结果结构
type PredictionResult struct {
	Value       float64 // 预测值
	Uncertainty float64 // 不确定性
}

// PredictionArray 预测数组结构
type PredictionArray struct {
	Results []PredictionResult // 预测结果数组
	Count   int                // 结果数量
}

// MpcCreatePredictor 创建 MPC 预测器
func MpcCreatePredictor(historicalData []float64) unsafe.Pointer {
	if len(historicalData) == 0 {
		return nil
	}

	// 将 Go 切片转换为 C 数组
	cData := make([]C.double, len(historicalData))
	for i, v := range historicalData {
		cData[i] = C.double(v)
	}

	return C.mpc_create_predictor(&cData[0], C.int(len(historicalData)))
}

// MpcPredict 执行预测
func MpcPredict(predictor unsafe.Pointer, predictionSteps int) *PredictionArray {
	if predictor == nil {
		return nil
	}

	cArray := C.mpc_predict(predictor, C.int(predictionSteps))
	if cArray == nil {
		return nil
	}

	// 转换 C 结构到 Go 结构
	count := int(cArray.count)
	results := make([]PredictionResult, count)

	for i := 0; i < count; i++ {
		cResult := (*C.PredictionResult)(unsafe.Pointer(uintptr(unsafe.Pointer(cArray.results)) + uintptr(i)*unsafe.Sizeof(C.PredictionResult{})))
		results[i] = PredictionResult{
			Value:       float64(cResult.value),
			Uncertainty: float64(cResult.uncertainty),
		}
	}

	return &PredictionArray{
		Results: results,
		Count:   count,
	}
}

// MpcUpdateMeasurement 更新测量值
func MpcUpdateMeasurement(predictor unsafe.Pointer, measurement float64) {
	if predictor != nil {
		C.mpc_update_measurement(predictor, C.double(measurement))
	}
}

// MpcGetCurrentState 获取当前状态
func MpcGetCurrentState(predictor unsafe.Pointer) PredictionResult {
	if predictor == nil {
		return PredictionResult{}
	}

	cResult := C.mpc_get_current_state(predictor)
	return PredictionResult{
		Value:       float64(cResult.value),
		Uncertainty: float64(cResult.uncertainty),
	}
}

// MpcFreePredictor 释放预测器
func MpcFreePredictor(predictor unsafe.Pointer) {
	if predictor != nil {
		C.mpc_free_predictor(predictor)
	}
}

// MpcFreePredictionArray 释放预测数组
func MpcFreePredictionArray(array *PredictionArray) {
	if array != nil {
		// 注意：这里我们不需要释放 Go 结构，因为内存管理由 Go 的 GC 处理
		// 但如果有 C 分配的内存，需要在这里释放
	}
}
