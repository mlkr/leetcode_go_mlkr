package problem551

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	ans bool
}{
	{
		"PPALLP",
		true,
	},

	{
		"PPALLPA",
		false,
	},

	{
		"PPALLL",
		false,
	},
}

func Test_checkRecord(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, checkRecord(q.s), "输入%v", q.s)
	}
}
func Benchmark_checkRecord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			checkRecord(q.s)
		}
	}
}
