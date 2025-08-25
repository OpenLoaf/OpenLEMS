package c_func

import (
	"fmt"
)

// ExampleAggregateAvgInt 演示平均值聚合函数的使用
func ExampleAggregateAvgInt() {
	// 整数平均值
	values := []any{1, 2, 3, 4, 5}
	avg, err := AggregateAvgInt(values)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Printf("平均值: %d\n", avg)

	// 浮点数平均值
	floatValues := []any{1.5, 2.5, 3.5, 4.5, 5.5}
	avgFloat, err := AggregateAvgFloat64(floatValues)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Printf("浮点数平均值: %.2f\n", avgFloat)

	// 字符串转换
	strValues := []any{"10", "20", "30", "40", "50"}
	avgStr, err := AggregateAvgInt(strValues)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Printf("字符串转换平均值: %d\n", avgStr)

	// Output:
	// 平均值: 3
	// 浮点数平均值: 3.50
	// 字符串转换平均值: 30
}

// ExampleAggregateSumInt 演示求和聚合函数的使用
func ExampleAggregateSumInt() {
	// 整数求和
	values := []any{1, 2, 3, 4, 5}
	sum, err := AggregateSumInt(values)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Printf("总和: %d\n", sum)

	// 浮点数求和
	floatValues := []any{1.5, 2.5, 3.5, 4.5, 5.5}
	sumFloat, err := AggregateSumFloat64(floatValues)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Printf("浮点数总和: %.2f\n", sumFloat)

	// Output:
	// 总和: 15
	// 浮点数总和: 17.50
}

// ExampleAggregateMaxInt 演示最大值聚合函数的使用
func ExampleAggregateMaxInt() {
	// 整数最大值
	values := []any{1, 5, 3, 9, 2}
	max, err := AggregateMaxInt(values)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Printf("最大值: %d\n", max)

	// 浮点数最大值
	floatValues := []any{1.5, 5.5, 3.5, 9.5, 2.5}
	maxFloat, err := AggregateMaxFloat64(floatValues)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Printf("浮点数最大值: %.2f\n", maxFloat)

	// Output:
	// 最大值: 9
	// 浮点数最大值: 9.50
}

// ExampleAggregateMinInt 演示最小值聚合函数的使用
func ExampleAggregateMinInt() {
	// 整数最小值
	values := []any{5, 1, 3, 9, 2}
	min, err := AggregateMinInt(values)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Printf("最小值: %d\n", min)

	// 浮点数最小值
	floatValues := []any{5.5, 1.5, 3.5, 9.5, 2.5}
	minFloat, err := AggregateMinFloat64(floatValues)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Printf("浮点数最小值: %.2f\n", minFloat)

	// Output:
	// 最小值: 1
	// 浮点数最小值: 1.50
}

// ExampleAggregateEqualInt 演示相等性聚合函数的使用
func ExampleAggregateEqualInt() {
	// 相等值
	equalValues := []any{5, 5, 5, 5}
	result, err := AggregateEqualInt(equalValues)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Printf("相等值结果: %d\n", result)

	// 不相等值
	unequalValues := []any{1, 2, 3, 4}
	_, err = AggregateEqualInt(unequalValues)
	if err != nil {
		fmt.Printf("不相等值错误: %v\n", err)
	}

	// 通用相等性判断
	strValues := []any{"test", "test", "test"}
	strResult, err := AggregateEqual(strValues)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Printf("字符串相等结果: %s\n", strResult)

	// Output:
	// 相等值结果: 5
	// 不相等值错误: values are not equal
	// 字符串相等结果: test
}
