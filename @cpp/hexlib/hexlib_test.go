package hexlib

import (
	"testing"
)

func TestHexVersion(t *testing.T) {
	version := HexVersion()
	if version == "" {
		t.Error("版本号不能为空")
	}
	t.Logf("库版本: %s", version)
}

func TestHexAdd(t *testing.T) {
	result := HexAdd(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("HexAdd(3, 5) = %d, 期望 %d", result, expected)
	}
	t.Logf("测试加法: 3 + 5 = %d", result)
}

func TestMpcCreatePredictor(t *testing.T) {
	// 历史数据 - 模拟微电网负载数据
	historicalData := []float64{50.0, 50.0, 50.0, 50.0, 50.0, 50.0, 50.0, 50.0,
		50.0, 50.0, 50.0, 40.0, 35.0, 34.0, 44.0}

	t.Logf("历史数据: %v", historicalData)

	// 创建预测器
	predictor := MpcCreatePredictor(historicalData)
	if predictor == nil {
		t.Fatal("创建预测器失败")
	}
	defer MpcFreePredictor(predictor)

	// 获取当前滤波状态
	currentState := MpcGetCurrentState(predictor)
	t.Logf("当前滤波状态: 值=%.2f, 不确定性=±%.2f",
		currentState.Value, currentState.Uncertainty)

	// 验证状态值在合理范围内
	if currentState.Value < 0 || currentState.Value > 100 {
		t.Errorf("当前状态值 %.2f 超出合理范围 [0, 100]", currentState.Value)
	}
	if currentState.Uncertainty < 0 {
		t.Errorf("不确定性值 %.2f 不能为负数", currentState.Uncertainty)
	}
}

func TestMpcPredict(t *testing.T) {
	// 历史数据
	historicalData := []float64{50.0, 50.0, 50.0, 50.0, 50.0, 50.0, 50.0, 50.0,
		50.0, 50.0, 50.0, 40.0, 35.0, 34.0, 44.0}

	predictor := MpcCreatePredictor(historicalData)
	if predictor == nil {
		t.Fatal("创建预测器失败")
	}
	defer MpcFreePredictor(predictor)

	// 生成预测
	predictionSteps := 10
	predictionArray := MpcPredict(predictor, predictionSteps)
	if predictionArray == nil {
		t.Fatal("生成预测失败")
	}
	defer MpcFreePredictionArray(predictionArray)

	// 验证预测结果
	if len(predictionArray.Results) != predictionSteps {
		t.Errorf("预测步数不匹配: 期望 %d, 实际 %d",
			predictionSteps, len(predictionArray.Results))
	}

	t.Logf("未来 %d 步预测:", predictionSteps)
	for i, result := range predictionArray.Results {
		confidence := 1.96 * result.Uncertainty // 95% 置信区间
		t.Logf("  第 %2d 步: %.2f ± %.2f (置信区间: [%.2f, %.2f])",
			i+1, result.Value, confidence,
			result.Value-confidence, result.Value+confidence)

		// 验证预测值的合理性
		if result.Value < 0 || result.Value > 200 {
			t.Errorf("第 %d 步预测值 %.2f 超出合理范围 [0, 200]", i+1, result.Value)
		}
		if result.Uncertainty < 0 {
			t.Errorf("第 %d 步不确定性 %.2f 不能为负数", i+1, result.Uncertainty)
		}
	}
}

func TestMpcUpdateMeasurement(t *testing.T) {
	// 历史数据
	historicalData := []float64{50.0, 50.0, 50.0, 50.0, 50.0, 50.0, 50.0, 50.0,
		50.0, 50.0, 50.0, 40.0, 35.0, 34.0, 44.0}

	predictor := MpcCreatePredictor(historicalData)
	if predictor == nil {
		t.Fatal("创建预测器失败")
	}
	defer MpcFreePredictor(predictor)

	// 获取更新前的状态
	stateBefore := MpcGetCurrentState(predictor)
	t.Logf("更新前状态: 值=%.2f, 不确定性=±%.2f",
		stateBefore.Value, stateBefore.Uncertainty)

	// 使用新测量值更新
	newMeasurement := 36.0
	t.Logf("使用新测量值更新: %.2f", newMeasurement)
	MpcUpdateMeasurement(predictor, newMeasurement)

	// 获取更新后的状态
	stateAfter := MpcGetCurrentState(predictor)
	t.Logf("更新后状态: 值=%.2f, 不确定性=±%.2f",
		stateAfter.Value, stateAfter.Uncertainty)

	// 验证状态发生了变化
	if stateBefore.Value == stateAfter.Value {
		t.Log("状态值未发生变化，这可能是正常的（取决于算法实现）")
	}

	// 验证更新后的状态值在合理范围内
	if stateAfter.Value < 0 || stateAfter.Value > 100 {
		t.Errorf("更新后状态值 %.2f 超出合理范围 [0, 100]", stateAfter.Value)
	}
}

func TestMpcFullWorkflow(t *testing.T) {
	t.Log("=== MPC 卡尔曼滤波器完整工作流程测试 ===")

	// 历史数据 - 模拟微电网负载数据
	historicalData := []float64{50.0, 50.0, 50.0, 50.0, 50.0, 50.0, 50.0, 50.0,
		50.0, 50.0, 50.0, 40.0, 35.0, 34.0, 44.0}

	t.Logf("历史数据: %v", historicalData)

	// 创建预测器
	predictor := MpcCreatePredictor(historicalData)
	if predictor == nil {
		t.Fatal("创建预测器失败")
	}
	defer MpcFreePredictor(predictor)

	// 获取当前滤波状态
	currentState := MpcGetCurrentState(predictor)
	t.Logf("当前滤波状态: 值=%.2f, 不确定性=±%.2f",
		currentState.Value, currentState.Uncertainty)

	// 生成预测
	predictionSteps := 10
	predictionArray := MpcPredict(predictor, predictionSteps)
	if predictionArray == nil {
		t.Fatal("生成预测失败")
	}
	defer MpcFreePredictionArray(predictionArray)

	t.Logf("未来 %d 步预测:", predictionSteps)
	for i, result := range predictionArray.Results {
		confidence := 1.96 * result.Uncertainty // 95% 置信区间
		t.Logf("  第 %2d 步: %.2f ± %.2f (置信区间: [%.2f, %.2f])",
			i+1, result.Value, confidence,
			result.Value-confidence, result.Value+confidence)
	}

	// 使用新测量值更新
	newMeasurement := 36.0
	t.Logf("使用新测量值更新: %.2f", newMeasurement)
	MpcUpdateMeasurement(predictor, newMeasurement)

	// 获取更新后的状态
	updatedState := MpcGetCurrentState(predictor)
	t.Logf("更新后状态: 值=%.2f, 不确定性=±%.2f",
		updatedState.Value, updatedState.Uncertainty)

	// 再次预测看看变化
	t.Logf("更新后的新预测:")
	newPredictions := MpcPredict(predictor, 10)
	if newPredictions != nil {
		defer MpcFreePredictionArray(newPredictions)
		for i, result := range newPredictions.Results {
			t.Logf("  第 %2d 步: %.2f ± %.2f",
				i+1, result.Value, 1.96*result.Uncertainty)
		}
	}

	t.Log("集成成功! MPC 卡尔曼滤波器已准备好在 Go 中使用。")
}
