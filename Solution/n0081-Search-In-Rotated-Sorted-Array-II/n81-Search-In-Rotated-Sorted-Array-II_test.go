package main

import (
	"fmt"
	"testing"
)

type question struct {
	arg args
	ans bool
}

type args struct {
	array  []int
	target int
}

func TestSearchII(t *testing.T) {
	que := []question{
		{
			arg: args{[]int{1}, 1},
			ans: true,
		},
		{
			arg: args{[]int{1, 0, 1, 1, 1}, 1},
			ans: true,
		},
		{
			arg: args{[]int{5, 6, 7, 8, 8, 8, 8, 9, 15, 1, 2, 3, 3, 5}, 4},
			ans: false,
		},
	}

	for i, q := range que {
		t.Run(fmt.Sprintf("Question %d", i), func(t *testing.T) {
			want := q.ans
			got := searchII(q.arg.array, q.arg.target)
			if want != got {
				t.Errorf("The answer is wrong, question: %v, should got: %v, actual got: %v\n", q.arg.array, want, got)
			}
		})
	}
}
