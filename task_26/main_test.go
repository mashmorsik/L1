package main

import "testing"

func TestIsUnique(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "unique_true", args: args{str: "star"}, want: true},
		{name: "doubles", args: args{str: "abcdEfgabc"}, want: false},
		{name: "doubles_different_cases", args: args{str: "abcdEfgG"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUnique(tt.args.str); got != tt.want {
				t.Errorf("IsUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
