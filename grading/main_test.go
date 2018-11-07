package main

import (
	"reflect"
	"testing"
)

func Test_gradingStudents(t *testing.T) {
	type args struct {
		grades []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{args: args{grades: []int32{73, 67, 38, 33}}, want: []int32{75, 67, 40, 33}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gradingStudents(tt.args.grades); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gradingStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
