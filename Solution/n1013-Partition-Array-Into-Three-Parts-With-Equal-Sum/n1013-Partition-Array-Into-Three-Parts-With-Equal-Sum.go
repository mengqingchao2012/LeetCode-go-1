package main

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, v := range A {
		sum += v
	}

	if sum%3 != 0 {
		return false
	}

	partSum := sum / 3
	left, right := 0, len(A)-1
	//注意初始化条件，如果都初始化为0，会出错，原因在于，
	leftSum, rightSum := A[left], A[right]
	for left+1 < right {
		if leftSum == partSum && rightSum == partSum {
			return true
		}

		if leftSum != partSum {
			left++ //注意由于初始化条件决定，要先执行自加
			leftSum += A[left]
		}

		if rightSum != partSum {
			right-- //注意由于初始化条件决定，要先执行自减
			rightSum += A[right]
		}
	}
	return false
}
