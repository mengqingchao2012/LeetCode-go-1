package main

//Time：O(n^2)，Space：O(1)
//选择排序：数组分为无序区和有序区，每次从无序区中选出最大或者最小的元素放入有序区中
func SelectSort(nums *[]int) {
	if len(*nums) == 0 {
		return
	}

	//将数组左边部分定义为有序区
	for i := 0; i < len(*nums); i++ {
		idx := i // idx用来存储无序区中选出的值，这里选的是最小值
		for j := i + 1; j < len(*nums); j++ {
			if (*nums)[j] < (*nums)[idx] {
				idx = j
			}
		}
		(*nums)[i], (*nums)[idx] = (*nums)[idx], (*nums)[i]
	}
	return
}

//func main() {
//	arr := []int{2, 7, 5, 0, 8, 6, 9, 9}
//	SelectSort(&arr)
//	fmt.Println(arr)
//}
