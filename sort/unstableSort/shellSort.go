package main

//Time：O(n^2)，Space：O(1)
//希尔排序是插入排序的改进版本，选择不同的步长其时间复杂度也不同，当步长是length/2时，
//时间复杂度是n^2，把步长改为1，就得到插入排序
func ShellSort(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	for gap := length >> 1; gap > 0; gap >>= 1 { //步长每次折半
		for i := gap; i < length; i++ {
			cur := nums[i]
			j := i - gap
			for j >= 0 && nums[j] > cur {
				nums[j+gap] = nums[j]
				j -= gap
			}
			nums[j+gap] = cur
		}
	}
}

//func main() {
//	arr := []int{2,3,1,3,2,4,6,7,9,2,19}
//	ShellSort(arr)
//	fmt.Println(arr)
//}
