package main

import "testing"

func Test_arrayManipulation(t *testing.T) {
	type args struct {
		n       int32
		queries [][]int32
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{n: 5, queries: [][]int32{
				{1, 2, 100},
				{2, 5, 100},
				{3, 4, 100},
			}},
			want: int64(200),
		},
		{
			args: args{n: 10, queries: [][]int32{
				{1, 5, 3},
				{4, 8, 7},
				{6, 9, 1},
			}},
			want: int64(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayManipulation(tt.args.n, tt.args.queries); got != tt.want {
				t.Errorf("arrayManipulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
