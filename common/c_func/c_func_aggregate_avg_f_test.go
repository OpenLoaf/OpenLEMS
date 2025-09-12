package c_func

import (
	"testing"
)

func TestAggregateAvgInt(t *testing.T) {
	tests := []struct {
		name    string
		values  []any
		want    *int
		wantErr bool
	}{
		{
			name:    "空切片",
			values:  []any{},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "单个值",
			values:  []any{10},
			want:    func() *int { v := 10; return &v }(),
			wantErr: false,
		},
		{
			name:    "多个值",
			values:  []any{1, 2, 3, 4, 5},
			want:    func() *int { v := 3; return &v }(),
			wantErr: false,
		},
		{
			name:    "字符串转换",
			values:  []any{"10", "20", "30"},
			want:    func() *int { v := 20; return &v }(),
			wantErr: false,
		},
		{
			name:    "混合类型",
			values:  []any{10, "20", 30.0},
			want:    func() *int { v := 20; return &v }(),
			wantErr: false,
		},
		{
			name:    "无效转换",
			values:  []any{1, "invalid", 3},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateAvgInt(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateAvgInt() error = %v, wantErr %+v", err, tt.wantErr)
				return
			}
			if tt.want == nil && got != nil {
				t.Errorf("AggregateAvgInt() = %v, want nil", got)
			} else if tt.want != nil && got == nil {
				t.Errorf("AggregateAvgInt() = nil, want %v", *tt.want)
			} else if tt.want != nil && got != nil && *got != *tt.want {
				t.Errorf("AggregateAvgInt() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestAggregateAvgFloat64(t *testing.T) {
	tests := []struct {
		name    string
		values  []any
		want    *float64
		wantErr bool
	}{
		{
			name:    "空切片",
			values:  []any{},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "单个值",
			values:  []any{10.5},
			want:    func() *float64 { v := 10.5; return &v }(),
			wantErr: false,
		},
		{
			name:    "多个值",
			values:  []any{1.0, 2.0, 3.0, 4.0, 5.0},
			want:    func() *float64 { v := 3.0; return &v }(),
			wantErr: false,
		},
		{
			name:    "带小数",
			values:  []any{1.5, 2.5, 3.5},
			want:    func() *float64 { v := 2.5; return &v }(),
			wantErr: false,
		},
		{
			name:    "字符串转换",
			values:  []any{"10.5", "20.5", "30.5"},
			want:    func() *float64 { v := 20.5; return &v }(),
			wantErr: false,
		},
		{
			name:    "混合类型",
			values:  []any{10, "20.5", 30.0},
			want:    func() *float64 { v := 20.166666666666668; return &v }(),
			wantErr: false,
		},
		{
			name:    "无效转换",
			values:  []any{1.0, "invalid", 3.0},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateAvgFloat64(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateAvgFloat64() error = %v, wantErr %+v", err, tt.wantErr)
				return
			}
			if tt.want == nil && got != nil {
				t.Errorf("AggregateAvgFloat64() = %v, want nil", got)
			} else if tt.want != nil && got == nil {
				t.Errorf("AggregateAvgFloat64() = nil, want %v", *tt.want)
			} else if tt.want != nil && got != nil && *got != *tt.want {
				t.Errorf("AggregateAvgFloat64() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestAggregateAvgInt64(t *testing.T) {
	tests := []struct {
		name    string
		values  []any
		want    *int64
		wantErr bool
	}{
		{
			name:    "大数值",
			values:  []any{int64(1000000000), int64(2000000000), int64(3000000000)},
			want:    func() *int64 { v := int64(2000000000); return &v }(),
			wantErr: false,
		},
		{
			name:    "字符串大数值",
			values:  []any{"1000000000", "2000000000", "3000000000"},
			want:    func() *int64 { v := int64(2000000000); return &v }(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateAvgInt64(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateAvgInt64() error = %v, wantErr %+v", err, tt.wantErr)
				return
			}
			if tt.want == nil && got != nil {
				t.Errorf("AggregateAvgInt64() = %v, want nil", got)
			} else if tt.want != nil && got == nil {
				t.Errorf("AggregateAvgInt64() = nil, want %v", *tt.want)
			} else if tt.want != nil && got != nil && *got != *tt.want {
				t.Errorf("AggregateAvgInt64() = %v, want %v", *got, *tt.want)
			}
		})
	}
}
