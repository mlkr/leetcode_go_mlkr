package problem514

type S struct {
	index, steps int
}

func findRotateSteps(ring string, key string) int {
	// 记录每个字母的位置
	m := make(map[rune][]int, 26)
	for i, r := range ring {
		m[r] = append(m[r], i)
	}

	ss := make([]S, 1, len(ring))
	ss[0] = S{0, 0}

	size := len(ring)
	for _, r := range key {
		indexs := m[r]
		temp := make([]S, len(indexs))

		for i, index := range indexs {
			s := fromTo(ss[0], index, size)

			for i := 1; i < len(ss); i++ {
				s.steps = min(s.steps, fromTo(ss[i], index, size).steps)
			}

			temp[i] = s
		}

		ss = temp
	}

	res := ss[0].steps
	for i := 1; i < len(ss); i++ {
		res = min(res, ss[i].steps)
	}

	return res

}

// s 到 index 的状态变化
func fromTo(s S, index, size int) S {
	a := min(s.index, index)
	b := max(s.index, index)

	return S{
		index: index,
		steps: s.steps + min(b-a, size+a-b) + 1, // +1 是因为要按下
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
