package problem446

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{2, 4, 6, 8, 10},
		7,
	},

	// 2 4 6 8 10
	// 2 4 6
	// 2 4 6 8
	// 4 6 8
	// 2 4 6 8 10
	// 4 6 8 10
	// 6 8 10
	// 2 6 10

	// 2 4 6 8 10	res=1
	// 4[2]=1
	// 6[4]=1 6[2]=2
	// 8[6]=1 8[4]=1 8[2]=3

	{
		[]int{1, 1, 1, 1},
		5,
	},

	// 1[0]=1
	// 2[0]=1+2		res=1
	// 3[0]=1+2+4	res=1+1+3

	{
		[]int{1, 2},
		0,
	},

	{
		[]int{0, 1<<31 - 1, 1<<32 - 1},
		0,
	},
}

func Test_numberOfArithmeticSlices(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, numberOfArithmeticSlices(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_numberOfArithmeticSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			numberOfArithmeticSlices(q.nums)
		}
	}
}
