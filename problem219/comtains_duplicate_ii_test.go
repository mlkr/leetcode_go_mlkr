package problem219

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	k    int
	ans  bool
}{
	{
		[]int{1, 2, 3, 1},
		3,
		true,
	},

	{
		[]int{1, 0, 1, 1},
		1,
		true,
	},

	{
		[]int{1, 2, 3, 1, 2, 3},
		2,
		false,
	},
}

func Test_containsNearbyDuplicate(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, containsNearbyDuplicate(q.nums, q.k))
	}
}
