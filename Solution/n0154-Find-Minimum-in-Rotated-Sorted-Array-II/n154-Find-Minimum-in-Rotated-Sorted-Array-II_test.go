package main

import (
	"fmt"
	"testing"
)

type problem struct {
	question []int
	answer   int
}

func TestFindMin(t *testing.T) {
	questions := []problem{
		{
			question: []int{1,3,5},
			answer:   1,
		},
		{
			question: []int{2,2,2,5,0,1},
			answer:   0,
		},
		{
			question: []int{3,3,1,3},
			answer:   1,
		},
		{
			question: []int{1,3,3},
			answer: 1,
		},
	}

	for i, v := range questions {
		t.Run(fmt.Sprintf("Question:%d ", i), func(t *testing.T) {
			want := v.answer
			got := findMin(v.question)
			if want != got {
				t.Errorf("The answer is wrong, question: %v, should got: %v, actual got: %v\n", v.question, want, got)
			}
		})
	}
}
