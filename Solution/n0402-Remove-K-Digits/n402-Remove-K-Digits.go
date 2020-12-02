package main

import ."LeetCode-go/utils"

func removeKdigits(num string, k int) string {
	size := Min(len(num), k) // 记录下最多能够删除的元素的个数
	finalSize := len(num) - size // 最终的结果中还剩余的元素位数
	stack := make([]rune, 0, size)

	for _, c := range num {
		for size > 0 && len(stack) > 0 && stack[len(stack) - 1] > c {
			stack = stack[: len(stack) - 1]
			size--
		}

		// 用于避免前缀 0 加入结果集，这里多加一个判断，即当前栈中元素为空
		// 且待加入的元素为 0 时，不能加入
		if c != '0' || len(stack) != 0 {
			stack = append(stack, c)
		}
	}

	// 对于数组本身满足升序条件时，stack[len(stack) - 1] > c 这个条件始终不会被触发，
	// 因此所有元素都会被加入到 stack 中，故需要进行裁剪，只保留前 finalSize 个
	if len(stack) > finalSize {
		stack = stack[:finalSize]
	}

	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}