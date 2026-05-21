package problem566

func matrixReshape(mat [][]int, r int, c int) [][]int {
	rows := len(mat)
	cols := len(mat[0])
	if r*c != rows*cols {
		return mat
	}

	res := make([][]int, r)
	count := 0
	for i := range res {
		res[i] = make([]int, c)
		for j := range res[i] {
			res[i][j] = mat[count/cols][count%cols]
			count++
		}
	}

	return res
}

// 更快
func matrixReshape2(mat [][]int, r int, c int) [][]int {
	rows := len(mat)
	cols := len(mat[0])

	if rows*cols != r*c {
		return mat
	}

	elems := make([]int, 0, r*c)
	for i := range mat {
		elems = append(elems, mat[i]...)
	}

	res := make([][]int, r)
	for i := range res {
		res[i] = elems[i*c : (i+1)*c]
	}

	return res
}
