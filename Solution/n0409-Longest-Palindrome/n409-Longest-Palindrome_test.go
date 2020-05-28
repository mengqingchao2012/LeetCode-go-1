package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{"ccc"},
			3,
		},
		{
			"test2",
			args{"cccbb"},
			5,
		},
		{
			"test3",
			args{"ac"},
			1,
		},
		{
			"test4",
			args{"abccccdd"},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
