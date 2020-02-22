package main

//冒泡排序：稳定排序算法
//Time：O(n^2)，Space:O(1)
func BubbleSort(nums *[]int) {
	length := len(*nums)
	if length == 0 {
		return
	}

	//注意边界的取值：end > 0, i < end
	for end := length - 1; end > 0; end-- {
		for i := 0; i < end; i++ {
			if (*nums)[i] > (*nums)[i+1] {
				(*nums)[i], (*nums)[i+1] = (*nums)[i+1], (*nums)[i]
			}
		}
	}
}

//优化1：使用一个标记符来检查该轮循环有没有发生元素交换，若没有，说明数组已经有序，可以提前返回
func BubbleSortByFlag(nums *[]int) {
	length := len(*nums)
	if length == 0 {
		return
	}

	for end := length - 1; end > 0; end-- {
		swapped := false
		for i := 0; i < end; i++ {
			if (*nums)[i] > (*nums)[i+1] {
				(*nums)[i], (*nums)[i+1] = (*nums)[i+1], (*nums)[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

//优化2：使用一个变量idx来记录最后一次发生交换的元素的下标，则下一次遍历时，将右边界更新为下标值
func BubbleSortByIdx(nums *[]int) {
	length := len(*nums)
	if length == 0 {
		return
	}

	var idx int
	for end := length - 1; end > 0; {
		idx = 0
		for i := 0; i < end; i++ {
			if (*nums)[i] > (*nums)[i+1] {
				(*nums)[i], (*nums)[i+1] = (*nums)[i+1], (*nums)[i]
				idx = i
			}
		}
		end = idx
	}
}

//func main() {
//	nums := []int{8, 4, 5, 7, 8}
//	BubbleSortByFlag(&nums)
//	fmt.Println(nums)
//}
