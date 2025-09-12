package c_func

import (
	"testing"
)

func TestAggregateSumInt(t *testing.T) {
	tests := []struct {
		name    string
		values  []any
		want    *int
		wantErr bool
	}{
		{
			name:    "空切片",
			values:  []any{},
			want:    func() *int { v := 0; return &v }(),
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
			want:    func() *int { v := 15; return &v }(),
			wantErr: false,
		},
		{
			name:    "字符串转换",
			values:  []any{"10", "20", "30"},
			want:    func() *int { v := 60; return &v }(),
			wantErr: false,
		},
		{
			name:    "混合类型",
			values:  []any{10, "20", 30.0},
			want:    func() *int { v := 60; return &v }(),
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
			got, err := AggregateSumInt(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateSumInt() error = %v, wantErr %+v", err, tt.wantErr)
				return
			}
			if tt.want == nil && got != nil {
				t.Errorf("AggregateSumInt() = %v, want nil", got)
			} else if tt.want != nil && got == nil {
				t.Errorf("AggregateSumInt() = nil, want %v", *tt.want)
			} else if tt.want != nil && got != nil && *got != *tt.want {
				t.Errorf("AggregateSumInt() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestAggregateSumFloat64(t *testing.T) {
	tests := []struct {
		name    string
		values  []any
		want    *float64
		wantErr bool
	}{
		{
			name:    "空切片",
			values:  []any{},
			want:    func() *float64 { v := 0.0; return &v }(),
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
			want:    func() *float64 { v := 15.0; return &v }(),
			wantErr: false,
		},
		{
			name:    "带小数",
			values:  []any{1.5, 2.5, 3.5},
			want:    func() *float64 { v := 7.5; return &v }(),
			wantErr: false,
		},
		{
			name:    "字符串转换",
			values:  []any{"10.5", "20.5", "30.5"},
			want:    func() *float64 { v := 61.5; return &v }(),
			wantErr: false,
		},
		{
			name:    "混合类型",
			values:  []any{10, "20.5", 30.0},
			want:    func() *float64 { v := 60.5; return &v }(),
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
			got, err := AggregateSumFloat64(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("AggregateSumFloat64() error = %v, wantErr %+v", err, tt.wantErr)
				return
			}
			if tt.want == nil && got != nil {
				t.Errorf("AggregateSumFloat64() = %v, want nil", got)
			} else if tt.want != nil && got == nil {
				t.Errorf("AggregateSumFloat64() = nil, want %v", *tt.want)
			} else if tt.want != nil && got != nil && *got != *tt.want {
				t.Errorf("AggregateSumFloat64() = %v, want %v", *got, *tt.want)
			}
		})
	}
}
