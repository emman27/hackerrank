package main

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		head *SinglyLinkedListNode
	}
	tests := []struct {
		name string
		args args
		want *SinglyLinkedListNode
	}{
		{args: args{head: &SinglyLinkedListNode{data: 1}}, want: &SinglyLinkedListNode{data: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
