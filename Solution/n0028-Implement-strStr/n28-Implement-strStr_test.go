package main

import "testing"

func TestStrStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				haystack: "",
				needle:   "",
			},
			0,
		},
		{
			"test2",
			args{
				haystack: "abc",
				needle:   "",
			},
			0,
		},
		{
			"test3",
			args{
				haystack: "",
				needle:   "a",
			},
			-1,
		},
		{
			"test4",
			args{
				haystack: "a",
				needle:   "a",
			},
			0,
		},
		{
			"test5",
			args{
				haystack: "hello",
				needle:   "ll",
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}