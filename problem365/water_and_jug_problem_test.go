package problem365

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	x, y, z int
	ans     bool
}{
	// {
	// 	13,
	// 	11,
	// 	1,
	// 	true,
	// },

	// {
	// 	1,
	// 	0,
	// 	0,
	// 	true,
	// },

	{
		0,
		2,
		1,
		false,
	},

	{
		34,
		5,
		6,
		true,
	},

	{
		3,
		5,
		4,
		true,
	},

	{
		3,
		5,
		1,
		true,
	},

	{
		3,
		5,
		2,
		true,
	},

	{
		3,
		5,
		3,
		true,
	},

	{
		3,
		5,
		8,
		true,
	},

	{
		3,
		5,
		9,
		false,
	},

	{
		2,
		6,
		5,
		false,
	},
}

func Test_canMeasureWater(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, canMeasureWater(q.x, q.y, q.z), "输入 %v %v %v", q.x, q.y, q.z)
	}
}

func Benchmark_canMeasureWater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			canMeasureWater(q.x, q.y, q.z)
		}
	}
}
