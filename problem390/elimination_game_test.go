package problem390

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var questions = []struct{
	n, ans int
}{
	{
		12,
		6,
	},

	{
		1,
		1,
	},

	{
		2,
		2,
	},

	// 1-9 与 1-8 结果相同
	// 1 2 3 4 5 6 7 8 9
	//   2   4   6   8
	//   2       6
	//           6
	{
		9,
		6,
	},


	// 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
	//   2   4   6   8   10    12    14    16  
	//   2       6       10          14
	//           6                   14
	//           6
	{
		16,
		6,
	},

	// 1 2 3 4 5 6 7 8 9 10 11 12
	//   2   4   6   8   10    12
	//   2       6       10   
	//           6  
	{
		12,
		6,
	},

	{
		10,
		8,
	},
}

func Test_lastRemaining(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, lastRemaining(q.n), "输入 %d", q.n)
	}
}

func Benchmark_lastRemaining(b *testing.B) {
	for i:=0;i<b.N;i++{
		for _, q := range questions {
			lastRemaining(q.n)
		}		
	}
}