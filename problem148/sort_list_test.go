package problem148

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para []int
	ans  []int
}{
	{
		[]int{-84, 142, 41, -17, -71, 170, 186, 183, -21, -76, 76, 10, 29, 81, 112, -39, -6, -43, 58, 41, 111, 33, 69, 97, -38, 82, -44, -7, 99, 135, 42, 150, 149, -21, -30, 164, 153, 92, 180, -61, 99, -81, 147, 109, 34, 98, 14, 178, 105, 5, 43, 46, 40, -37, 23, 16, 123, -53, 34, 192, -73, 94, 39, 96, 115, 88, -31, -96, 106, 131, 64, 189, -91, -34, -56, -22, 105, 104, 22, -31, -43, 90, 96, 65, -85, 184, 85, 90, 118, 152, -31, 161, 22, 104, -85, 160, 120, -31, 144, 115},
		[]int{-96, -91, -85, -85, -84, -81, -76, -73, -71, -61, -56, -53, -44, -43, -43, -39, -38, -37, -34, -31, -31, -31, -31, -30, -22, -21, -21, -17, -7, -6, 5, 10, 14, 16, 22, 22, 23, 29, 33, 34, 34, 39, 40, 41, 41, 42, 43, 46, 58, 64, 65, 69, 76, 81, 82, 85, 88, 90, 90, 92, 94, 96, 96, 97, 98, 99, 99, 104, 104, 105, 105, 106, 109, 111, 112, 115, 115, 118, 120, 123, 131, 135, 142, 144, 147, 149, 150, 152, 153, 160, 161, 164, 170, 178, 180, 183, 184, 186, 189, 192},
	},

	{
		[]int{3, 4, 1},
		[]int{1, 3, 4},
	},

	{
		[]int{4, 2, 1, 3},
		[]int{1, 2, 3, 4},
	},

	{
		[]int{-1, 5, 3, 4, 0},
		[]int{-1, 0, 3, 4, 5},
	},
}

func Test_sortList(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		head := sliceToList(q.para)
		head = sortList(head)
		ast.Equal(q.ans, listToSlice(head))
	}
}

func Test_sortList2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		head := sliceToList(q.para)
		head = sortList2(head)
		ast.Equal(q.ans, listToSlice(head))
	}
}

func Test_sortList3(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		head := sliceToList(q.para)
		head = sortList3(head)
		ast.Equal(q.ans, listToSlice(head))
	}
}

func Benchmark_sortList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			head := sliceToList(q.para)
			head = sortList(head)
		}
	}
}

func Benchmark_sortList2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			head := sliceToList(q.para)
			head = sortList2(head)
		}
	}
}

func Benchmark_sortList3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			head := sliceToList(q.para)
			head = sortList3(head)
		}
	}
}
