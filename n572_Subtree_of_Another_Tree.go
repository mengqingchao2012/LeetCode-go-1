package main

import (
	. "LeetCode-go/utils"
	"hash/crc32"
	"strconv"
)


//遍历s的每个节点，与t进行比较；
//Time：O(m*n)，Space：O(h)，m和n是s和t各自的节点数，h是s的树高
func isSubtree(s *TreeNode, t *TreeNode) bool {
	var res bool
	switch {
	case t == nil: //空树是任意树的子树，所以t=nil时返回true
		res = true
	case s == nil: //s为空，则返回false，注意前两个case的顺序不能颠倒，因为t和s都为空时结果是true
		res = false
	default: //s和t都非空，则判断当前节点的值，递归判断t和s的左右子树的关系，任意一个满足就返回true
		res = isSameTrees(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
	}
	return res
}

func isSameTrees(s, t *TreeNode) bool {
	var res bool
	switch {
	case s == nil && t == nil:
		res = true
	case s == nil || t == nil:
		res = false
	default:
		res = s.Val == t.Val && isSameTrees(s.Left, t.Left) && isSameTrees(s.Right, t.Right)
	}
	return res
}

//思路二：借助hash函数，为每颗子树生成一个hash值，这样一来，就不需要每次都遍历t子树，只需比较其hash值即可
//Time：O(m + n)，Space：O(m + n)，m和n是s和t各自的节点数
var mp = make(map[*TreeNode]uint32) //用来保存每个子树的唯一生成hash值

func isSubtree1(s *TreeNode, t *TreeNode) bool {
	computeHash(s)
	computeHash(t)
	return isSub(s, t)
}

func isSub(s , t *TreeNode) bool {
	var res bool
	switch {
	case t == nil:
		res = true
	case s == nil:
		res = false
	default:
		res = mp[s] == mp[t] || isSubtree(s.Left, t) || isSubtree(s.Right, t)
	}
	return res
}

//递归的遍历树的节点，以左子树的节点值+根节点值+右子树的节点值构成字符串，然后再使用hash函数转成hash值
func computeHash(t *TreeNode) string {
	if t == nil {
		return "#"
	}

	left := computeHash(t.Left)
	right := computeHash(t.Right)
	hash := left + strconv.Itoa(t.Val) + right
	mp[t] = crc32.ChecksumIEEE([]byte(hash))
	return hash
}
