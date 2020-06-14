package problem393

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	data []int
	ans  bool
}{
	{
		[]int{240, 162, 138, 147},
		true,
	},

	{
		[]int{237},
		false,
	},

	{
		[]int{},
		false,
	},

	{
		[]int{197, 130, 1},
		true,
	},

	{
		[]int{235, 140, 4},
		false,
	},

	{
		[]int{235, 140, 141},
		true,
	},

	{
		[]int{235, 140, 141, 127},
		true,
	},

	{
		[]int{245, 140, 141, 142},
		true,
	},

	{
		[]int{245, 140, 141, 142, 1},
		true,
	},

	{
		[]int{245, 140, 141, 1},
		false,
	},

	{
		[]int{127, 127},
		true,
	},

	{
		[]int{192},
		false,
	},

	{
		[]int{240},
		false,
	},

	{
		[]int{192, 127},
		false,
	},

	{
		[]int{192, 128},
		true,
	},

	{
		[]int{250},
		false,
	},
}

func Test_validUtf8(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, validUtf8(q.data), "输入 %v", q.data)
	}
}

func Benchmark_validUtf8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			validUtf8(q.data)
		}
	}
}
