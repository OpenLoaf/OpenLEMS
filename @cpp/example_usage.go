package main

import (
	"fmt"

	"hexlib"
)

func main() {
	fmt.Println("MPC 卡尔曼滤波器 - C++ 到 Go 集成示例")
	fmt.Println("===========================================")

	// 测试基础函数
	fmt.Printf("库版本: %s\n", hexlib.HexVersion())
	fmt.Printf("测试加法: %d + %d = %d\n", 3, 5, hexlib.HexAdd(3, 5))

	// 历史数据 - 模拟微电网负载数据
	historicalData := []float64{50.0, 60.0, 80.0, 54.0, 55.0, 70.0, 88.0, 65.0,
		50.0, 50.0, 50.0, 40.0, 35.0, 34.0, 35.0}

	fmt.Printf("\n历史数据: %v\n", historicalData)

	// 创建预测器
	predictor := hexlib.MpcCreatePredictor(historicalData)
	if predictor == nil {
		fmt.Println("创建预测器失败")
		return
	}
	defer hexlib.MpcFreePredictor(predictor)

	// 获取当前滤波状态
	currentState := hexlib.MpcGetCurrentState(predictor)
	fmt.Printf("\n当前滤波状态:\n")
	fmt.Printf("  值: %.2f\n", currentState.Value)
	fmt.Printf("  不确定性: ±%.2f\n", currentState.Uncertainty)

	// 生成预测
	predictionSteps := 10
	predictionArray := hexlib.MpcPredict(predictor, predictionSteps)
	if predictionArray == nil {
		fmt.Println("生成预测失败")
		return
	}

	fmt.Printf("\n未来 %d 步预测:\n", predictionSteps)
	for i, result := range predictionArray.Results {
		confidence := 1.96 * result.Uncertainty // 95% 置信区间
		fmt.Printf("  第 %2d 步: %.2f ± %.2f (置信区间: [%.2f, %.2f])\n",
			i+1, result.Value, confidence,
			result.Value-confidence, result.Value+confidence)
	}

	// 使用新测量值更新
	newMeasurement := 36.0
	fmt.Printf("\n使用新测量值更新: %.2f\n", newMeasurement)
	hexlib.MpcUpdateMeasurement(predictor, newMeasurement)

	// 获取更新后的状态
	updatedState := hexlib.MpcGetCurrentState(predictor)
	fmt.Printf("更新后状态:\n")
	fmt.Printf("  值: %.2f\n", updatedState.Value)
	fmt.Printf("  不确定性: ±%.2f\n", updatedState.Uncertainty)

	// 再次预测看看变化
	fmt.Printf("\n更新后的新预测:\n")
	newPredictions := hexlib.MpcPredict(predictor, 5)
	if newPredictions != nil {
		for i, result := range newPredictions.Results {
			fmt.Printf("  第 %2d 步: %.2f ± %.2f\n",
				i+1, result.Value, 1.96*result.Uncertainty)
		}
	}

	fmt.Println("\n集成成功! MPC 卡尔曼滤波器已准备好在 Go 中使用。")
}
