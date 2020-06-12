package problem391

type point struct {
	x, y int
}

func isRectangleCover(rectangles [][]int) bool {
	if len(rectangles) == 0 {
		return false
	}

	if len(rectangles[0]) == 0 {
		return false
	}

	m := make(map[point]bool)

	minX, minY := rectangles[0][0], rectangles[0][1]
	maxX, maxY := rectangles[0][2], rectangles[0][3]

	area := 0

	for _, rect := range rectangles {
		minX = min(minX, rect[0])
		minY = min(minY, rect[1])

		maxX = max(maxX, rect[2])
		maxY = max(maxY, rect[3])

		area += (rect[2] - rect[0]) * (rect[3] - rect[1])

		for _, x := range []int{rect[0], rect[2]} {
			for _, y := range []int{rect[1], rect[3]} {
				p := point{x, y}
				if m[p] {
					delete(m, p)
				} else {
					m[p] = true
				}
			}
		}
	}

	for _, x := range []int{minX, maxX} {
		for _, y := range []int{minY, maxY} {
			p := point{x, y}
			if !m[p] {
				return false
			}
		}
	}

	if len(m) != 4 ||
		area != (maxX-minX)*(maxY-minY) {
		return false
	}

	return true

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
