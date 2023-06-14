package problem514

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	ring, key string
	ans       int
}{
	{
		"godding",
		"gd",
		4,
	},

	{
		"godding",
		"godding",
		13,
	},
}

func Test_findRotateSteps(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findRotateSteps(q.ring, q.key), "输入 %s %s", q.ring, q.key)
	}
}

func Benchmark_findRotateSteps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findRotateSteps(q.ring, q.key)
		}
	}
}
