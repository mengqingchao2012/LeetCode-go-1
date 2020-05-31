package main

// 其实就是求斐波那契数列的第n项，Time：O(n)，Space：O(1)
// 要注意初始条件：f(0) = 1, f(1) = 1, f(2) = 2
func climbStairs(n int) int {
	first, second := 1, 1

	for i := 1; i < n; i++ { // 注意边界是 i < n
		first, second = second, first+second
	}
	return second
}
