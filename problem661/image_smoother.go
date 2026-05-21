package problem661

func imageSmoother(img [][]int) [][]int {
	rows := len(img)
	cols := len(img[0])

	res := make([][]int, rows)
	for i := range res {
		res[i] = make([]int, cols)
	}

	for i := range img {
		for j := range img[i] {
			count := 0
			sum := 0

			for r := i - 1; r <= i+1; r++ {
				if r < 0 {
					continue
				}

				if r == rows {
					break
				}

				for c := j - 1; c <= j+1; c++ {
					if c < 0 || c == cols {
						continue
					}

					sum += img[r][c]
					count++

				}
			}

			res[i][j] = sum / count
		}

	}

	return res
}
