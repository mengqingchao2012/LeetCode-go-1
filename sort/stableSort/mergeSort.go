package main

//Time：O(nlogn)，Space：O(n)
func merge(nums *[]int, low, mid, high int) {
	i, j := low, mid+1
	var temp []int
	for i <= mid && j <= high {
		if (*nums)[i] <= (*nums)[j] {
			temp = append(temp, (*nums)[i])
			i++
		} else {
			temp = append(temp, (*nums)[j])
			j++
		}
	}

	if i <= mid {
		temp = append(temp, (*nums)[i:]...)
	}

	if j <= mid {
		temp = append(temp, (*nums)[j:]...)
	}

	copy((*nums)[low:high+1], temp)
}

func mergeSort(nums *[]int, low, high int) {
	if low < high {
		mid := low + ((high - low) >> 1)
		mergeSort(nums, low, mid)
		mergeSort(nums, mid+1, high)
		merge(nums, low, mid, high)
	}
}

func SortRecursive(arr []int) {
	if len(arr) == 0 {
		return
	}
	mergeSort(&arr, 0, len(arr)-1)
}

//func main() {
//	arr := []int{5, 4, 3, 2, 1, 1}
//	SortRecursive(arr)
//	fmt.Println(arr)
//}
