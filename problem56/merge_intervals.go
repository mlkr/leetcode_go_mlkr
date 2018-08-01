package problem56

import "math/rand"

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	len := len(intervals)
	if len <= 1 {
		return intervals
	}

	quickSort(intervals)

	temp := intervals[0]
	res := make([]Interval, 0, len)
	for i := 1; i < len; i++ {
		if temp.End >= intervals[i].Start {
			temp.End = max(temp.End, intervals[i].End)
			continue
		}

		res = append(res, temp)
		temp = intervals[i]
	}

	res = append(res, temp)

	return res
}

func quickSort(intervals []Interval) {
	len := len(intervals)
	if len <= 1 {
		return
	}

	j := rand.Intn(len)
	intervals[0], intervals[j] = intervals[j], intervals[0]

	j = partition(intervals)
	quickSort(intervals[0:j])
	quickSort(intervals[j+1:])
}

func partition(intervals []Interval) int {
	len := len(intervals)
	i, j := 1, len-1
	for {
		for intervals[0].Start >= intervals[i].Start && i < len-1 {
			i++
		}

		for intervals[0].Start <= intervals[j].Start && j > 0 {
			j--
		}

		if i >= j {
			break
		}

		intervals[i], intervals[j] = intervals[j], intervals[i]
	}

	intervals[0], intervals[j] = intervals[j], intervals[0]

	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
