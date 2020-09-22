package main

import (
	"fmt"
	"testing"
)
import ."LeetCode-go/utils"

func Test(t *testing.T) {
	head := &ListNode{Val: 3}
	h1 := &ListNode{Val: 5}
	head.Next = h1

	res := reverseBetween(head, 1, 2)
	for res.Next != nil {
		fmt.Printf("res.Val: %v, res.Next: %v\n", res.Val, res.Next)
		res = res.Next
	}

}
