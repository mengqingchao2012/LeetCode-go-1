package main

import "math"

//前序遍历的特点：根节点-左子树节点-右子树节点，其结果数组可以分成三部分
//根节点 小于根节点的数（左子树节点） 大于根节点的数（右子树节点）
//方法一：递归求解：Time：O(n^2)，Space：O(n)
func verifyPreorder(preorder []int) bool {
	//本题认定空数组也满足二叉搜索树的前序遍历
	return verify(&preorder, 0, len(preorder))
}

//验证某个子序列是否满足二叉搜索树前序遍历的特点，取值区间【start, end)
func verify(preorder *[]int, start, end int) bool {
	//start == end 说明数组中没有元素，start + 1 == end 说明数组中只有一个元素，均返回true
	if start == end || start + 1 == end {
		return true
	}
	root := (*preorder)[start] //取出根节点
	i := start + 1 //找到左子树的开头
	for ; i < end && (*preorder)[i] < root; {
		i++
	}
	mid := i //找到右子树的开头
	for ; i < end && (*preorder)[i] > root; {
		i++
	}
	if i == end { //注意判断条件 i == end 不能漏
		//分别判断左子树和右子树是否满足要求，注意判断范围start 和 end 的取值
		return verify(preorder, start + 1, mid) && verify(preorder, mid, end)
	}
	return false
}

//方法二：使用一个辅助栈
//二叉树的特点:从根节点开始沿着左子树一直向下，数值一直递减，当数值开始增大时，说明已经不在原来的左子树支路上了,
//而是换到了右子树
//使用一个辅助栈，遍历数组，当数组的值持续递减时（stack[top] < v），将值依次入栈；
//当遇到一个比栈顶元素大的值时，说明已经从左子树转换到了右子树上，则将栈中的元素依次出栈，同时用出栈元素更新lowerBound，
//这样lowerBound的值就是子树的根节点的值，也就是小于右子树元素的最大值；
//如果新加入的元素比lowerBound还小，则不满足前序遍历的特点，返回false；
//当遍历完整个数组后，如果没有提前返回，就说明满足条件，返回true
func verifyPreorder1(preorder []int) bool {
	if len(preorder) == 0 {
		return true
	}

	stact := make([]int, len(preorder))
	top := -1
	lowerBound := math.MinInt64 //注意lowerBound初始化的值
	for _, v := range preorder {
		if v < lowerBound {
			return false
		}
		for ;top != -1 && stact[top] < v; {
			lowerBound = stact[top]
			top--
		}
		top++
		stact[top] = v
	}
	return true
}
