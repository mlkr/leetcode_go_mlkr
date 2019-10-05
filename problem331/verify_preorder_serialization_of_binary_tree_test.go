package problem331

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	preorder string
	ans      bool
}{
	{
		"#,7,6,9,#,#,#",
		false,
	},

	{
		"#",
		true,
	},

	{
		"9,3,4,#,#,1,#,#,2,#,6,#,#",
		true,
	},

	{
		"9,3,4,#,#,1,#,#,2,#,#,#,#",
		false,
	},

	{
		"1,#",
		false,
	},

	{
		"9,#,#,1",
		false,
	},

	{
		"9,#,#,#,1",
		false,
	},
}

func Test_isValidSerialization(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, isValidSerialization(q.preorder), "输入 %v", q.preorder)
	}
}

func Benchmark_isValidSerialization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isValidSerialization(q.preorder)
		}
	}
}
