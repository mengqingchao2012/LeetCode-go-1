package main

//Time: O(n*log(n)), Space: O(1)
//由于堆是完全二叉树，所以用数组保存时，满足：
//父节点下标是i，则左孩子的下标是 2i+1，右孩子的下标是 2i+2
//孩子节点的下标是i，则父节点的下标是 (i-1)/2
//对于完全二叉树来说，下标从 n/2 + 1 到 n 的节点都是叶子节点
func HeapSort(nums *[]int) {
	length := len(*nums)
	if length == 0 {
		return
	}

	buildHeap(nums, length-1)         //将数组转为最大堆
	for i := length - 1; i > 0; i-- { //注意边界：i > 0
		//每次将堆顶元素置换到最后一个叶子节点的位置，该元素就完成排序
		(*nums)[0], (*nums)[i] = (*nums)[i], (*nums)[0]
		heapify(nums, 0, i-1) // 重新堆化满足最大堆的要求
	}
}

//将数组转化为堆，注意建堆的时间复杂度是O(n)
func buildHeap(nums *[]int, end int) {
	for i := end / 2; i >= 0; i-- { //注意边界：i>=0，因为要遍历到根节点
		heapify(nums, i, end)
	}
}

// Time: O(log(n))
func heapify(nums *[]int, i, end int) {
	parent := i
	child := 2*parent + 1
	for child <= end {
		//从左右孩子中挑选出最大的一个
		if child+1 <= end && (*nums)[child+1] > (*nums)[child] {
			child++
		}
		//如果此时父节点的值大于最大孩子的值，直接break
		if (*nums)[parent] >= (*nums)[child] {
			break
		}
		//否则就交换父节点和最大的孩子
		(*nums)[parent], (*nums)[child] = (*nums)[child], (*nums)[parent]
		//同时交换以后如果后续还有孩子节点，还要接着比较
		parent = child
		child = 2*parent + 1
	}
}

//func main() {
//	arr := []int{2,3,1,3,2,4,6,7,9,2,19}
//	HeapSort(&arr)
//	fmt.Println(arr)
//}
