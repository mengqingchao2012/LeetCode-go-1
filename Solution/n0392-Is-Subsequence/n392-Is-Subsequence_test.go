package main

import (
	"fmt"
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{"", "ahbgdc"},
			want: true,
		},
		{
			args: args{"acb", "ahbgdc"},
			want: false,
		},
		{
			args: args{"adc", "ahbgdc"},
			want: true,
		},
		{
			args: args{"axb", "ahbgdc"},
			want: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Question:%d ", i), func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}