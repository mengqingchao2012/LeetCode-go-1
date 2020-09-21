package main

import (
	"strconv"
	"testing"
)

func Test_firstMissingPositive(t *testing.T) {
	tests := []struct {
		args []int
		want int
	}{
		{
			args: []int{1, 1},
			want: 2,
		},
		{
			args: []int{},
			want: 1,
		},
		{
			args: []int{2, 1},
			want: 3,
		},
		{
			args: []int{7,8,9,11,12},
			want: 1,
		},
		{
			args: []int{3,4,-1,1},
			want: 2,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := firstMissingPositive(tt.args); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
