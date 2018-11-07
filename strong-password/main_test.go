package main

import "testing"

func Test_minimumNumber(t *testing.T) {
	type args struct {
		n    int32
		pass string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{args: args{pass: "Ab1"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumNumber(tt.args.n, tt.args.pass); got != tt.want {
				t.Errorf("minimumNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
