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
			question: []int{3, 4, 5, 1, 2},
			answer:   1,
		},
		{
			question: []int{4, 5, 6, 7, 0, 1, 2},
			answer:   0,
		},
		{
			question: []int{1},
			answer:   1,
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
