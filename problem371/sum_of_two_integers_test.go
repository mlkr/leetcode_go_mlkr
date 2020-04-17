package problem371

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	a, b, ans int
}{
	{
		1,
		2,
		3,
	},

	{
		-2,
		3,
		1,
	},

	{255, 255, 510},
	//   1111 1111
	//   1111 1111
	// 1 1111 1110 (记录进位 按位与 向左进一位)
	//   0000 0000 (不考虑进位 半加)
}

func Test_getSum(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, getSum(q.a, q.b), "输入 %q %q", q.a, q.b)
	}
}

func Benchmark_getSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			getSum(q.a, q.b)
		}
	}
}
