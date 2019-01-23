package main

import "testing"

func Test_minimumSwaps(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{args: args{arr: []int32{4, 3, 1, 2}}, want: 3},
		{args: args{arr: []int32{1, 3, 5, 2, 4, 6, 7}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSwaps(tt.args.arr); got != tt.want {
				t.Errorf("minimumSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
