package problem442

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct{
	nums, ans []int
}{
	{
		[]int{4,3,2,7,8,2,3,1},
		[]int{2,3},
	},
}

func Test_findDuplicates(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findDuplicates(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_findDuplicates(b *testing.B) {
	for i:=0;i<b.N;i++{
		for _, q := range questions {
			findDuplicates(q.nums)
		}
	}
}