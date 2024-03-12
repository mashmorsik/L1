package main

import (
	"reflect"
	"testing"
)

func TestSortTemp(t *testing.T) {
	type args struct {
		temp []float32
	}
	tests := []struct {
		name string
		args args
		want map[int][]float32
	}{
		{name: "below_zero_above_zero_same_values",
			args: args{temp: []float32{-10.2, 10.2, -15.0, 15.0}},
			want: map[int][]float32{-10: {-10.2, -15}, 10: {10.2, 15}}},
		{name: "zero_value_test",
			args: args{temp: []float32{-1.0, 1.0}},
			want: map[int][]float32{0: {-1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortTemp(tt.args.temp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortTemp() = %v, want %v", got, tt.want)
			}
		})
	}
}
