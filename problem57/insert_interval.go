package problem57

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

func insert(intervals []Interval, newInterval Interval) []Interval {
	len := len(intervals)
	if len == 0 {
		return []Interval{newInterval}
	}

	if intervals[0].Start > newInterval.End {
		return append([]Interval{newInterval}, intervals...)
	}

	if intervals[len-1].End < newInterval.Start {
		return append(intervals, newInterval)
	}

	res := make([]Interval, 0, len)
	for i := range intervals {
		if inRange(intervals[i], newInterval) {
			newInterval = merge(intervals[i], newInterval)
			if i == len-1 {
				res = append(res, newInterval)
				break
			}
			continue
		}

		if newInterval.Start > intervals[i].End {
			res = append(res, intervals[i])
			continue
		}

		if newInterval.End < intervals[i].Start {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			break
		}

	}

	return res
}

func inRange(a, b Interval) bool {
	return (a.Start <= b.Start && b.Start <= a.End) || (b.Start <= a.Start && a.Start <= b.End)
}

func merge(a, b Interval) Interval {
	return Interval{
		Start: min(a.Start, b.Start),
		End:   max(a.End, b.End),
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
