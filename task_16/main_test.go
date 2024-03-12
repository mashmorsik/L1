package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "one_element", args: args{nums: []int{3}}, want: []int{3}},
		{name: "empty", args: args{nums: []int{}}, want: []int{}},
		{name: "negatives", args: args{nums: []int{-5, -2, -10, -3}}, want: []int{-10, -5, -3, -2}},
		{name: "positives", args: args{nums: []int{3, 6, 1, 2}}, want: []int{1, 2, 3, 6}},
		{name: "positives_and_negatives", args: args{nums: []int{3, -5, 7, -2}}, want: []int{-5, -2, 3, 7}},
		{name: "with_zero", args: args{nums: []int{7, 3, 0, -1}}, want: []int{-1, 0, 3, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
