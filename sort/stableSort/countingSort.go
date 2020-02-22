package main

var maxRange = 1000 //对于数组中的元素 x，有 0 <= x <= maxRange

//Time：O(n + k)，Space：O(n + k)，n是元素个数，k是范围maxRange
func CountingSort(nums []int) {
	if len(nums) == 0 {
		return
	}

	count := make([]int, maxRange+1)
	res := make([]int, maxRange+1)

	//遍历数组，统计各个数字出现的次数
	for _, v := range nums {
		count[v]++
	}

	//遍历count结果集，针对每个下标（即数组中对应的数字），存储其起始位置的下标
	start := 0
	for i, v := range count {
		count[i] = start
		start += v //下一个数字的起始位置等于start加上当前数字出现的次数
	}

	//再次遍历数组
	for _, v := range nums {
		idx := count[v] //从count结果集中取出该元素的插入下标
		res[idx] = v    //更新结果集
		count[v]++      //并将该元素的插入下标后移
	}
	copy(nums, res)
}

//func main() {
//	arr := []int{2,3,1,3,2,4,6,7,9,2,19}
//	CountingSort(arr)
//	fmt.Println(arr)
//}
