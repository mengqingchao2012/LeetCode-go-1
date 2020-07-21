package main

import (
	"LeetCode-go/utils"
	"sort"
)

type Intervals struct {
	start int
	end   int
}

//Time：O(nlogn)，Space：O(1)
func merge1(intervals []Intervals) []Intervals {
	size := len(intervals)
	if size == 0 {
		return intervals
	}

	res := make([]Intervals, 0, size)

	sort.Slice(intervals, func(i, j int) bool { //注意此处对slice排序的写法
		return intervals[i].start < intervals[j].start
	})

	for _, v := range intervals {
		//如果当前结果是第一组数，或者结果集中最后一组数的end小于当前数组的start，说明当前数组和结果集中的最后一组数没有包含
		//关系，则当前数组可以直接添加进结果集
		n := len(res)
		if n == 0 || res[n-1].end < v.start {
			res = append(res, v)
		} else { //否则，说明结果集中的最后一组数组与当前数组有包含关系，使用end元素较大的那个值更新结果集
			res[n-1].end = utils.Max(res[n-1].end, v.end)
		}
	}
	return res
}

//不使用结构体的写法
func mergeForSlice(intervals [][]int) [][]int {
	size := len(intervals)
	if size == 0 {
		return intervals
	}

	res := make([][]int, 0, size)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for _, v := range intervals {
		n := len(res)
		if n == 0 || res[n-1][1] < v[0] {
			res = append(res, v)
		} else {
			res[n-1][1] = utils.Max(res[n-1][1], v[1])
		}
	}
	return res
}
