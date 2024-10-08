package problem593

import "sort"

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	distances := []int{
		getDistance(p1, p2),
		getDistance(p1, p3),
		getDistance(p1, p4),
		getDistance(p2, p3),
		getDistance(p2, p4),
		getDistance(p3, p4),
	}

	sort.Ints(distances)

	return distances[0] > 0 &&
		distances[0] == distances[1] &&
		distances[1] == distances[2] &&
		distances[2] == distances[3] &&
		distances[4] == distances[5]
}

func getDistance(p1, p2 []int) int {
	x := p1[0] - p2[0]
	y := p1[1] - p2[1]

	return x*x + y*y
}
