package main

import (
	"reflect"
	"testing"
)

func Test_freqQuery(t *testing.T) {
	type args struct {
		queries [][]int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{args: args{queries: [][]int32{{1, 1}, {2, 2}, {3, 2}, {1, 1}, {1, 1}, {2, 1}, {3, 2}}}, want: []int32{0, 1}},
		{args: args{queries: [][]int32{{1, 5}, {1, 6}, {3, 2}, {1, 10}, {1, 10}, {1, 6}, {2, 5}, {3, 2}}}, want: []int32{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := freqQuery(tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("freqQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
