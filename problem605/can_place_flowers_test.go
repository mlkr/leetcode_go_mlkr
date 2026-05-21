package problem605

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	flowerbed []int
	n         int
	ans       bool
}{

	{
		[]int{0, 0, 0, 0},
		2,
		true,
	},

	{
		[]int{1},
		1,
		false,
	},

	{
		[]int{0, 0, 1, 0, 1},
		1,
		true,
	},

	{
		[]int{1, 0, 0, 0, 1},
		1,
		true,
	},

	{
		[]int{1, 0, 0, 0, 1},
		2,
		false,
	},
}

func Test_canPlaceFlowers(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, canPlaceFlowers(q.flowerbed, q.n))
	}
}

func Benchmark_canPlaceFlowers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			canPlaceFlowers(q.flowerbed, q.n)
		}

	}
}
