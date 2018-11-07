package main

import "testing"

func Test_timeConversion(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "07:45:45AM"}, want: "07:45:45"},
		{args: args{s: "07:45:45PM"}, want: "19:45:45"},
		{args: args{s: "12:45:45PM"}, want: "12:45:45"},
		{args: args{s: "12:45:45AM"}, want: "00:45:45"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeConversion(tt.args.s); got != tt.want {
				t.Errorf("timeConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}
