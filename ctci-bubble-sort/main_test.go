package main

import "testing"

func Test_countSwapsInternal(t *testing.T) {
	type args struct {
		a []int32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{a: []int32{1, 2, 3}}, want: 0},
		{args: args{a: []int32{3, 2, 1}}, want: 3},
		{args: args{a: []int32{4, 2, 3, 1}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSwapsInternal(tt.args.a); got != tt.want {
				t.Errorf("countSwapsInternal() = %v, want %v", got, tt.want)
			}
		})
	}
}
