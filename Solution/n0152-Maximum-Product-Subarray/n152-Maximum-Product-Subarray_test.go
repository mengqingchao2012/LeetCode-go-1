package main

import (
	"fmt"
	"testing"
)

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{ []int{2,3,-2,4} },
			want: 6,
		},
		{
			args: args{ []int{0, 2} },
			want: 2,
		},
		{
			args: args{ []int{8, 1, -2, 4, -1} },
			want: 64,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Question %d", i), func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
