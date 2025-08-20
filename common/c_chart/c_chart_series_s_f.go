package c_chart

// Series 表示图表数据系列
type Series struct {
	Name string   `json:"name"`
	Type string   `json:"type"`
	Unit string   `json:"unit"`
	Data []string `json:"data"`
}

// AppendData 向系列添加数据
func (s *Series) AppendData(value string) {
	s.Data = append(s.Data, value)
}

// NewSeries 创建新的数据系列
func NewSeries(name string, seriesType string, unit string, capacity int) *Series {
	return &Series{
		Name: name,
		Type: seriesType,
		Unit: unit,
		Data: make([]string, 0, capacity),
	}
}
