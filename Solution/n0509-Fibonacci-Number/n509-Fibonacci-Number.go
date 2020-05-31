package main

func fib(N int) int {
	if N == 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	first, second := 0, 1
	for i := 2; i <= N; i++ {
		first, second = second, first+second
	}
	return second
}
