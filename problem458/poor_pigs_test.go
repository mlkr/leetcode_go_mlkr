package problem458

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	buckets, minutesToDie, minutesToTest int
	ans                                  int
}{
	{
		1,
		1,
		1,
		0,
	},

	{
		1000,
		15,
		60,
		5,
	},

	{
		25, 15, 60, 2,
	},
}

func Test_poorPigs(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, poorPigs(q.buckets, q.minutesToDie, q.minutesToTest), "buckets %v", q.buckets)
	}
}

func Benchmark_poorPigs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			poorPigs(q.buckets, q.minutesToDie, q.minutesToTest)
		}
	}
}
