package problem587

func outerTrees(trees [][]int) [][]int {
	size := len(trees)
	if size < 4 {
		return trees
	}

	// 选取左下角的点作为开始的点
	first := 0
	for i := 1; i < size; i++ {
		if trees[first][0] == trees[i][0] &&
			trees[first][1] > trees[i][1] {
			first = i
		}

		if trees[first][0] > trees[i][0] {
			first = i
		}
	}

	res := make([][]int, 0, size)
	res = append(res, trees[first])

	m := make(map[int]bool, size)
	m[first] = true

	// 以cur点为原点，顺时针方向扫描
	// 最后的点next和cur点的连线为一段护栏
	// next点为作为新的cur点，再次扫描
	// 如果护栏的点cur first重合，扫描完毕，围栏完成
	cur := first
	for {
		next := (cur + 1) % size

		for i := 0; i < size; i++ {
			if (m[i] && i != first) || i == next {
				continue
			}

			res_s := scan(trees[cur], trees[next], trees[i])
			if res_s < 0 {
				next = i
			}

			if res_s == 0 {
				if isFurther(trees[cur], trees[next], trees[i]) {
					next = i
				}
			}
		}

		// 将cur-next连线上的点都加入护栏
		for i := 0; i < size; i++ {
			if m[i] {
				continue
			}

			if scan(trees[cur], trees[next], trees[i]) == 0 {
				res = append(res, trees[i])
				m[i] = true
			}
		}

		cur = next
		if cur == first {
			break
		}
	}

	return res
}

// o点为原点，线oa的斜率为 ay/ax, 线ob的斜率为 by/bx
// 如果线oa ob 线率相等，则点a b 在一条直线上 ay/ax = by/bx => ax*by = ay*bx => ax*by - ay*bx = 0
// ax*by - ay*bx = 0，点a b 在一条直线上

// ax*by - ay*bx < 0 => by/bx < ay/ax (ax bx 同号) or by/bx > ay/ax（(ax bx 异号)）
// 通过画直角坐标系可得 o点为原点顺时针扫描先到a再到b
// ax*by - ay*bx < 0，o点为原点顺时针扫描先到a再到b

// 注意：角aob 的角度最大为180度，因为如果大于180度，线ab可以作为一段护栏将o围在里面，而所有选取的o点都是边缘点，不可能在围栏里。
func scan(o, a, b []int) int {
	return (a[0]-o[0])*(b[1]-o[1]) - (a[1]-o[1])*(b[0]-o[0])
}

// b点比a点更远？
func isFurther(o, a, b []int) bool {
	return distance(o, b)-distance(o, a) > 0
}

func distance(o, a []int) int {
	return (a[0]-o[0])*(a[0]-o[0]) + (a[1]-o[1])*(a[1]-o[1])
}
