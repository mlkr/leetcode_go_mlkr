package problem497

import "math/rand"

type Solution struct {
	scopes []int
	rects  [][]int
}

func Constructor(rects [][]int) Solution {
	s := Solution{}
	s.rects = rects
	s.scopes = make([]int, 0, len(rects))

	total := 0
	for _, rect := range rects {
		if len(rect) != 0 {
			total += (rect[3] - rect[1] + 1) * (rect[2] - rect[0] + 1) // +1 是因为 [a,b] 之间一共有 b-a+1 个离散的点
			s.scopes = append(s.scopes, total)
		}
	}

	return s
}

func (this *Solution) Pick() []int {
	size := len(this.scopes)
	if size == 0 {
		return nil
	}

	r := rand.Intn(this.scopes[size-1] + 1)

	i := 0
	for ; i < size; i++ {
		if this.scopes[i] >= r {
			break
		}
	}

	rect := this.rects[i]
	x := rect[0] + rand.Intn((rect[2] - rect[0] + 1))
	y := rect[1] + rand.Intn((rect[3] - rect[1] + 1))

	return []int{x, y}
}
