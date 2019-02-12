package main

import "testing"

func Test_maximumToys(t *testing.T) {
	type args struct {
		prices []int32
		k      int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{args: args{prices: []int32{1, 12, 5, 111, 200, 1000, 10}, k: 50}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumToys(tt.args.prices, tt.args.k); got != tt.want {
				t.Errorf("maximumToys() = %v, want %v", got, tt.want)
			}
		})
	}
}
