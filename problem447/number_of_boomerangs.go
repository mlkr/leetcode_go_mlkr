package problem447

func numberOfBoomerangs(points [][]int) int {
	size := len(points)
	if size < 3 {
		return 0
	}

	res := 0
	for i := 0; i < size; i++ {
		m := make(map[int]int, size)
		for j := 0; j < size; j++ {
			if i == j {
				continue
			}

			d := square(points[i], points[j])
			m[d]++
		}

		for _, v := range m {
			res += v * (v - 1)
		}
	}

	return res
}

func square(p1, p2 []int) int {
	x := p1[0] - p2[0]
	y := p1[1] - p2[1]

	return x*x + y*y
}
