package main

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "value_present_in_the_slice_positive_int",
			args: args{nums: []int{9, 5, 3, 1, 2}, n: 5},
			want: "The index of 5 in the sorted array is 3"},
		{name: "value_present_in_the_slice_negative_int",
			args: args{nums: []int{-9, 5, 3, 1, 2}, n: -9},
			want: "The index of -9 in the sorted array is 0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.nums, tt.args.n); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
