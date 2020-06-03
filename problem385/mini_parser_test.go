package problem385

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	ans *NestedInteger
}{
	{
		"",
		nil,
	},

	{
		"324",
		&NestedInteger{num: 324},
	},

	{
		"[123,[456,[789]]]",
		&NestedInteger{
			ns: []*NestedInteger{
				&NestedInteger{num: 123},
				&NestedInteger{
					ns: []*NestedInteger{
						&NestedInteger{num: 456},
						&NestedInteger{
							ns: []*NestedInteger{
								&NestedInteger{num: 789},
							},
						},
					},
				},
			},
		},
	},
}

func Test_deserialize(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, deserialize(q.s), "输入 %s", q.s)
	}
}

func Benchmark_deserialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			deserialize(q.s)
		}
	}
}
