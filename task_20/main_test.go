package main

import "testing"

func TestFlipperStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "English", args: args{str: "Not my cup of tea"}, want: "tea of cup my Not"},
		{name: "Russian", args: args{str: "Самый странный подарок"}, want: "подарок странный Самый"},
		{name: "Chinese", args: args{str: "有 意 思"}, want: "思 意 有"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlipperStr(tt.args.str); got != tt.want {
				t.Errorf("FlipperStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFlipperStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FlipperStr("Not my cup of tea")
	}
}

func BenchmarkFlipperStrBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FlipperStrBuilder("Not my cup of tea")
	}
}
