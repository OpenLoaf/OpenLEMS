package c_func

import (
	"testing"
)

func TestAggregateMaxInt(t *testing.T) {
	tests := []struct {
		name    string
		values  []any
		want    int
		wantErr bool
	}{
		{
			name:    "空切片",
			values:  []any{},
			want:    0,
			wantErr: false,
		},
		{
			name:    "单个值",
			values:  []any{10},
			want:    10,
			wantErr: false,
		},
		{
			name:    "多个值",
			values:  []any{1, 2, 3, 4, 5},
			want:    5,
			wantErr: false,
		},
		{
			name:    "字符串转换",
			values:  []any{"10", "20", "30"},
			want:    30,
			wantErr: false,
		},
		{
			name:    "混合类型",
			values:  []any{10, "20", 30.0},
			want:    30,
			wantErr: false,
		},
		{
			name:    "负数",
			values:  []any{-1, -2, -3},
			want:    -1,
			wantErr: false,
		},
		{
			name:    "无效转换",
			values:  []any{1, "invalid", 3},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateMaxInt(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateMaxInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AggregateMaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregateMaxFloat64(t *testing.T) {
	tests := []struct {
		name    string
		values  []any
		want    float64
		wantErr bool
	}{
		{
			name:    "空切片",
			values:  []any{},
			want:    0.0,
			wantErr: false,
		},
		{
			name:    "单个值",
			values:  []any{10.5},
			want:    10.5,
			wantErr: false,
		},
		{
			name:    "多个值",
			values:  []any{1.0, 2.0, 3.0, 4.0, 5.0},
			want:    5.0,
			wantErr: false,
		},
		{
			name:    "带小数",
			values:  []any{1.5, 2.5, 3.5},
			want:    3.5,
			wantErr: false,
		},
		{
			name:    "字符串转换",
			values:  []any{"10.5", "20.5", "30.5"},
			want:    30.5,
			wantErr: false,
		},
		{
			name:    "混合类型",
			values:  []any{10, "20.5", 30.0},
			want:    30.0,
			wantErr: false,
		},
		{
			name:    "负数",
			values:  []any{-1.5, -2.5, -3.5},
			want:    -1.5,
			wantErr: false,
		},
		{
			name:    "无效转换",
			values:  []any{1.0, "invalid", 3.0},
			want:    0.0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateMaxFloat64(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateMaxFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AggregateMaxFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregateMaxInt64(t *testing.T) {
	tests := []struct {
		name    string
		values  []any
		want    int64
		wantErr bool
	}{
		{
			name:    "大数值",
			values:  []any{int64(1000000000), int64(2000000000), int64(3000000000)},
			want:    3000000000,
			wantErr: false,
		},
		{
			name:    "字符串大数值",
			values:  []any{"1000000000", "2000000000", "3000000000"},
			want:    3000000000,
			wantErr: false,
		},
		{
			name:    "负数大数值",
			values:  []any{int64(-1000000000), int64(-2000000000), int64(-3000000000)},
			want:    -1000000000,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateMaxInt64(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateMaxInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AggregateMaxInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregateMaxUint(t *testing.T) {
	tests := []struct {
		name    string
		values  []any
		want    uint
		wantErr bool
	}{
		{
			name:    "空切片",
			values:  []any{},
			want:    0,
			wantErr: false,
		},
		{
			name:    "单个值",
			values:  []any{uint(10)},
			want:    10,
			wantErr: false,
		},
		{
			name:    "多个值",
			values:  []any{uint(1), uint(2), uint(3), uint(4), uint(5)},
			want:    5,
			wantErr: false,
		},
		{
			name:    "字符串转换",
			values:  []any{"10", "20", "30"},
			want:    30,
			wantErr: false,
		},
		{
			name:    "大数值",
			values:  []any{uint(1000000000), uint(2000000000), uint(3000000000)},
			want:    3000000000,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateMaxUint(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateMaxUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AggregateMaxUint() = %v, want %v", got, tt.want)
			}
		})
	}
}
