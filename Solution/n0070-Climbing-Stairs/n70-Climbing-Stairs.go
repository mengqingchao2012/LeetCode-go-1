package main

//其实就是求斐波那契数列的第n项，Time：O(n)，Space：O(1)
func climbStairs(n int) int {
	first, second := 1, 1

	for i := 1; i < n; i++ {
		first, second = second, first+second
	}
	return second
}
