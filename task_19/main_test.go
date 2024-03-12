package main

import "testing"

func TestFlipper(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "English", args: args{str: "AvadaKedavra"}, want: "arvadeKadavA"},
		{name: "Russian", args: args{str: "главрыба"}, want: "абырвалг"},
		{name: "Chinese", args: args{str: "什么事"}, want: "事么什"},
		{name: "emoji", args: args{str: "🤘 ☄ ★ 💪"}, want: "💪 ★ ☄ 🤘"},
		{name: "emoji2", args: args{str: "😈 👿"}, want: "👿 😈"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Flipper(tt.args.str); got != tt.want {
				t.Errorf("Flipper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFlipper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Flipper("главрыба")
	}
}

func BenchmarkFlipperRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FlipperRune("главрыба")
	}
}

func BenchmarkFlipperByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FlipperByte("главрыба")
	}
}
