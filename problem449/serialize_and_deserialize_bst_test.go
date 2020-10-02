package problem449

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

var questions = []struct {
	pre []int
}{
	{
		[]int{8, 3, 1, 6, 4, 7, 10, 14, 13},
	},

	{
		nil,
	},

	{
		[]int{-1},
	},
}

func Test_serialize(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		nodes := []rune{}
		for _, p := range q.pre {
			nodes = append(nodes, rune(p))
		}

		obj := Constructor()
		root := obj.deserialize(string(nodes))
		ast.Equal(string(nodes), obj.serialize(root), "输入 %v", nodes)
	}

}

func Benchmark_serialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			nodes := []rune{}
			for _, p := range q.pre {
				nodes = append(nodes, rune(p))
			}

			obj := Constructor()
			root := obj.deserialize(string(nodes))
			obj.serialize(root)
		}

	}
}
