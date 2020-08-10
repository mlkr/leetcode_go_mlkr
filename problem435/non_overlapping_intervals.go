package problem435

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int)bool{
		return intervals[i][1] < intervals[j][1]
	})

	end := -1<<31
	res := 0
	for _, v := range intervals{
		if end <= v[0] {
			end = v[1]
		}else{
			res++
		}
	}

	return res
}