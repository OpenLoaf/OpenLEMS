package c_func

import (
	"reflect"
	"testing"
)

func TestAggregateEqual(t *testing.T) {
	tests := []struct {
		name    string
		values  []any
		want    any
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
			want:    10,
			wantErr: false,
		},
		{
			name:    "相等值",
			values:  []any{5, 5, 5, 5},
			want:    5,
			wantErr: false,
		},
		{
			name:    "字符串相等",
			values:  []any{"test", "test", "test"},
			want:    "test",
			wantErr: false,
		},
		{
			name:    "混合类型相等",
			values:  []any{10, "10", 10.0},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "不相等值",
			values:  []any{1, 2, 3},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "字符串不相等",
			values:  []any{"a", "b", "c"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "复杂类型相等",
			values:  []any{[]int{1, 2}, []int{1, 2}, []int{1, 2}},
			want:    []int{1, 2},
			wantErr: false,
		},
		{
			name:    "复杂类型不相等",
			values:  []any{[]int{1, 2}, []int{1, 3}, []int{1, 2}},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateEqual(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateEqual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AggregateEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregateEqualInt(t *testing.T) {
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
			name:    "相等值",
			values:  []any{5, 5, 5, 5},
			want:    5,
			wantErr: false,
		},
		{
			name:    "字符串转换相等",
			values:  []any{"10", "10", "10"},
			want:    10,
			wantErr: false,
		},
		{
			name:    "混合类型相等",
			values:  []any{10, "10", 10.0},
			want:    10,
			wantErr: false,
		},
		{
			name:    "不相等值",
			values:  []any{1, 2, 3},
			want:    0,
			wantErr: true,
		},
		{
			name:    "无效转换",
			values:  []any{1, "invalid", 1},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateEqualInt(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateEqualInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AggregateEqualInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregateEqualFloat64(t *testing.T) {
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
			name:    "相等值",
			values:  []any{5.5, 5.5, 5.5},
			want:    5.5,
			wantErr: false,
		},
		{
			name:    "字符串转换相等",
			values:  []any{"10.5", "10.5", "10.5"},
			want:    10.5,
			wantErr: false,
		},
		{
			name:    "混合类型相等",
			values:  []any{10.5, "10.5", 10.5},
			want:    10.5,
			wantErr: false,
		},
		{
			name:    "不相等值",
			values:  []any{1.0, 2.0, 3.0},
			want:    0.0,
			wantErr: true,
		},
		{
			name:    "无效转换",
			values:  []any{1.0, "invalid", 1.0},
			want:    0.0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateEqualFloat64(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateEqualFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AggregateEqualFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregateEqualInt64(t *testing.T) {
	tests := []struct {
		name    string
		values  []any
		want    int64
		wantErr bool
	}{
		{
			name:    "大数值相等",
			values:  []any{int64(1000000000), int64(1000000000), int64(1000000000)},
			want:    1000000000,
			wantErr: false,
		},
		{
			name:    "字符串大数值相等",
			values:  []any{"1000000000", "1000000000", "1000000000"},
			want:    1000000000,
			wantErr: false,
		},
		{
			name:    "大数值不相等",
			values:  []any{int64(1000000000), int64(2000000000), int64(1000000000)},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateEqualInt64(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateEqualInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AggregateEqualInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
