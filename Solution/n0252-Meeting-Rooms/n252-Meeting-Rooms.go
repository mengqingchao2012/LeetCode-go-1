package main

import "sort"

func canAttendMeetings(intervals [][]int) bool {
	n := len(intervals)
	if n == 0 {
		return true
	}

	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	for i := 0; i < n - 1; i++ {
		if intervals[i][1] > intervals[i + 1][0] {
			return false
		}
	}
	return true
}
