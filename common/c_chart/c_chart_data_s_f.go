package c_chart

import "fmt"

// ChartData 表示完整的图表数据
type ChartData struct {
	XAxis  XAxis    `json:"xAxis"`
	Series []Series `json:"series"`
}

// AddTimestamp 向X轴添加时间戳
func (c *ChartData) AddTimestamp(timestamp int64) {
	c.XAxis.Data = append(c.XAxis.Data, fmt.Sprintf("%d", timestamp))
}

// AddSeries 添加数据系列
func (c *ChartData) AddSeries(series Series) {
	c.Series = append(c.Series, series)
}

// NewChartData 创建新的图表数据
func NewChartData(pointCount int) *ChartData {
	return &ChartData{
		XAxis: XAxis{
			Type: ChartTypeCategory,
			Data: make([]string, 0),
		},
		Series: make([]Series, 0, pointCount),
	}
}
