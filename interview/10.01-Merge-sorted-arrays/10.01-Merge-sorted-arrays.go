package main

func merge(A []int, m int, B []int, n int)  {
	l := m + n - 1
	n -= 1
	m -= 1
	for m >= 0 || n >= 0{
		switch {
		case m < 0:
			A[l] = B[n]
			n--
		case n < 0:
			A[l] = A[m]
			m--
		case A[m] > B[n]:
			A[l] = A[m]
			m--
		default:
			A[l] = B[n]
			n--
		}
		l--
	}
}
