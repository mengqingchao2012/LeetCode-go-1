package main

//Time：O(nlogn)，Space：O(n)
func mergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + ((right - left) >> 1)
	mergeSort(nums, left, mid)
	mergeSort(nums, mid + 1, right)

	i, j := left, mid + 1
	tmp := make([]int, 0, right - left + 1)
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}

	if i <= mid {
		tmp = append(tmp, nums[i : mid + 1]...)
	}
	if j <= right {
		tmp = append(tmp, nums[j : right + 1]...)
	}
	copy(nums[left : right + 1], tmp)
}

func SortRecursive(arr []int) {
	if len(arr) == 0 {
		return
	}
	mergeSort(arr, 0, len(arr)-1)
}

//func main() {
//	arr := []int{5, 4, 3, 2, 1, 1}
//	SortRecursive(arr)
//	fmt.Println(arr)
//}
