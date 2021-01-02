package main

func maxSlidingWindow(nums []int, k int) []int {
	deque := make([]int, 0, k) // 单调队列（双端队列），保存元素的下标
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if len(deque) != 0 && i - k + 1 > deque[0] { // 如果队头元素已经超出了窗口范围，则弹出队头
			deque = deque[1:]
		}
		for len(deque) != 0 && nums[i] >= nums[deque[len(deque) - 1]] { // 如果队尾元素对应的值小于等于当前元素的值，则弹出队尾
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i) // 将当前元素入队
		if i >= k - 1 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}
