package problem149

type Point struct {
	X int
	Y int
}

func maxPoints(points []Point) int {
	n := len(points)
	// diffMap 用来过滤掉相同的点，并记录他们的个数
	diffMap := make(map[Point]int, n)
	for i := 0; i < n; i++ {
		diffMap[points[i]]++
	}

	size := len(diffMap)
	if size <= 2 {
		// 最多只有两个不同的点, 则所有点都在一条直线上
		return n
	}

	if n > size {
		// 去除 points 中的重复点
		points = make([]Point, 0, size)
		for point := range diffMap {
			points = append(points, point)
		}
	}

	max := 0
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {

			count := 0
			for k := 0; k < size; k++ {
				if isSameLine(points[i], points[j], points[k]) {
					count += diffMap[points[k]]
				}
			}

			if count > max {
				max = count
			}
		}
	}

	return max
}

func isSameLine(p1, p2, p3 Point) bool {
	return (p3.X-p1.X)*(p2.Y-p1.Y) == (p2.X-p1.X)*(p3.Y-p1.Y)
}
