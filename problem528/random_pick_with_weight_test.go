package problem528

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution(t *testing.T) {
	ast := assert.New(t)

	w := []int{1, 3}
	total := 0
	for i := range w {
		total += w[i]
	}

	picks := 1000
	count := make([]int, len(w))
	s := Constructor(w)
	for i := 0; i < picks; i++ {
		count[s.PickIndex()]++
	}

	for i, c := range count {
		exp := float64(w[i]) / float64(total)
		act := float64(c) / float64(picks)
		ast.InDelta(exp, act, 0.001)
	}
}
