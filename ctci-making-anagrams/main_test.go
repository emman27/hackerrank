package main

import "testing"

func Test_makeAnagram(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"zero", args{a: "abc", b: "abc"}, 0},
		{"simple", args{a: "abcc", b: "abc"}, 1},
		{"newchar", args{a: "abcd", b: "abc"}, 1},
		{"newchar in b", args{a: "abc", b: "abce"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeAnagram(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("makeAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
