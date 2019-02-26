package main

import "testing"

func Test_alternatingCharacters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"zero", args{"A"}, 0},
		{"one", args{"AA"}, 1},
		{"correct", args{"AB"}, 0},
		{"one, but not first one", args{"ABB"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alternatingCharacters(tt.args.s); got != tt.want {
				t.Errorf("alternatingCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
