package main

import (
	"fmt"
	"testing"
)

type problem struct {
	question []int
	answer   int
}

func TestMissingNumber (t *testing.T) {
	questions := []problem{
		{
			question: []int{0, 1, 2, 3, 4, 5, 6, 7, 9},
			answer:   8,
		},
		{
			question: []int{0},
			answer:   1,
		},
		{
			question: []int{0, 1},
			answer:   2,
		},
		{
			question: []int{1},
			answer:   0,
		},
	}

	for i, v := range questions {
		t.Run(fmt.Sprintf("Question:%d ", i), func(t *testing.T) {
			want := v.answer
			got := missingNumber(v.question)
			if want != got {
				t.Errorf("The answer is wrong, question: %v, should got: %v, actual got: %v\n", v.question, want, got)
			}
		})
	}
}
