package c_func

import (
	"reflect"
	"testing"
)

func TestAggregateEqualInt(t *testing.T) {
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
			wantErr: true,
		},
		{
			name:    "单个值",
			values:  []any{10},
			want:    func() *int { v := 10; return &v }(),
			wantErr: false,
		},
		{
			name:    "相等值",
			values:  []any{5, 5, 5, 5},
			want:    func() *int { v := 5; return &v }(),
			wantErr: false,
		},
		{
			name:    "字符串转换相等",
			values:  []any{"10", "10", "10"},
			want:    func() *int { v := 10; return &v }(),
			wantErr: false,
		},
		{
			name:    "混合类型相等",
			values:  []any{10, "10", 10.0},
			want:    func() *int { v := 10; return &v }(),
			wantErr: false,
		},
		{
			name:    "不相等值",
			values:  []any{1, 2, 3},
			want:    nil,
			wantErr: true,
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
			got, err := AggregateEqualInt(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateEqualInt() error = %v, wantErr %+v", err, tt.wantErr)
				return
			}
			if tt.want == nil && got != nil {
				t.Errorf("AggregateEqualInt() = %v, want nil", got)
			} else if tt.want != nil && got == nil {
				t.Errorf("AggregateEqualInt() = nil, want %v", *tt.want)
			} else if tt.want != nil && got != nil && *got != *tt.want {
				t.Errorf("AggregateEqualInt() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestAggregateEqualFloat64(t *testing.T) {
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
			wantErr: true,
		},
		{
			name:    "单个值",
			values:  []any{10.5},
			want:    func() *float64 { v := 10.5; return &v }(),
			wantErr: false,
		},
		{
			name:    "相等值",
			values:  []any{5.5, 5.5, 5.5, 5.5},
			want:    func() *float64 { v := 5.5; return &v }(),
			wantErr: false,
		},
		{
			name:    "字符串转换相等",
			values:  []any{"10.5", "10.5", "10.5"},
			want:    func() *float64 { v := 10.5; return &v }(),
			wantErr: false,
		},
		{
			name:    "混合类型相等",
			values:  []any{10.0, "10.0", 10.0},
			want:    func() *float64 { v := 10.0; return &v }(),
			wantErr: false,
		},
		{
			name:    "不相等值",
			values:  []any{1.0, 2.0, 3.0},
			want:    nil,
			wantErr: true,
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
			got, err := AggregateEqualFloat64(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateEqualFloat64() error = %v, wantErr %+v", err, tt.wantErr)
				return
			}
			if tt.want == nil && got != nil {
				t.Errorf("AggregateEqualFloat64() = %v, want nil", got)
			} else if tt.want != nil && got == nil {
				t.Errorf("AggregateEqualFloat64() = nil, want %v", *tt.want)
			} else if tt.want != nil && got != nil && *got != *tt.want {
				t.Errorf("AggregateEqualFloat64() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

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
			wantErr: true,
		},
		{
			name:    "单个值",
			values:  []any{"hello"},
			want:    "hello",
			wantErr: false,
		},
		{
			name:    "相等字符串",
			values:  []any{"hello", "hello", "hello"},
			want:    "hello",
			wantErr: false,
		},
		{
			name:    "相等结构体",
			values:  []any{struct{ A int }{1}, struct{ A int }{1}},
			want:    struct{ A int }{1},
			wantErr: false,
		},
		{
			name:    "不相等值",
			values:  []any{"hello", "world"},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AggregateEqual(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateEqual() error = %v, wantErr %+v", err, tt.wantErr)
				return
			}
			if tt.want == nil && got != nil {
				t.Errorf("AggregateEqual() = %v, want nil", got)
			} else if tt.want != nil && got == nil {
				t.Errorf("AggregateEqual() = nil, want %v", tt.want)
			} else if tt.want != nil && got != nil {
				// 使用 DeepEqual 比较 any 类型
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("AggregateEqual() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
