package problem452

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	end := points[0][1]
	res := 1
	for i := range points {
		if end >= points[i][0] {
			continue
		}

		res++
		end = points[i][1]
	}

	return res
}
