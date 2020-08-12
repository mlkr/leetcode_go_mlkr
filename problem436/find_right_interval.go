package problem436

import "sort"

type Interval struct{
	start, index int
}

func findRightInterval(intervals [][]int) []int {
	size := len(intervals)
	isSort := make([]Interval, size)
	for i, v := range intervals {
		isSort[i] = Interval{v[0], i} 
	}

	sort.Slice(isSort, func(i, j int) bool {
		return isSort[i].start < isSort[j].start
	})

	res := make([]int, size)
	for i, v := range intervals {
		r := sort.Search(size, func(i int) bool {
			return isSort[i].start >= v[1]
		})

		res[i] = -1
		if r < size {
			res[i] = isSort[r].index
		}
	}
   

	return res
}
