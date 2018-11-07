package main

import (
	"reflect"
	"testing"
)

func Test_getTotalX(t *testing.T) {
	type args struct {
		a []int32
		b []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{args: args{a: []int32{2, 4}, b: []int32{16, 32, 96}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTotalX(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getTotalX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCommonFactors(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{args: args{arr: []int32{2, 4}}, want: []int32{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCommonFactors(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCommonFactors() = %v, want %v", got, tt.want)
			}
		})
	}
}
