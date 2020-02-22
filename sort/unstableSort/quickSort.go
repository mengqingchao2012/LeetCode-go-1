package main

//Time：O(nlogn)，Space：O(n)
func lomutoPartition(nums *[]int, start, end int) int {
	pivot := (*nums)[end]
	i := start
	for j := start; j < end; j++ {
		if (*nums)[j] < pivot {
			(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
			i++
		}
	}
	(*nums)[i], (*nums)[end] = (*nums)[end], (*nums)[i]
	return i
}

func lomutoSort(nums *[]int, start, end int) {
	if start < end {
		//注意 lomuto 分隔得到的分隔点 k 的位置就是其最终位置
		k := lomutoPartition(nums, start, end)
		lomutoSort(nums, start, k-1)
		lomutoSort(nums, k+1, end)
	}
}

func LomutoSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	lomutoSort(&nums, 0, len(nums)-1)
}

//func main() {
//	arr := []int{2, 7, 5, 0, 8, 6, 9, 9}
//	LomutoSort(arr)
//	fmt.Println(arr)
//}
