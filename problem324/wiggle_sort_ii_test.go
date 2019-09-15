package problem324

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = [][]int{
	[]int{1, 5, 1, 1, 6, 4},
	[]int{1, 3, 2, 2, 3, 1},
	[]int{1, 1, 2, 1, 2, 2, 1},
	[]int{4, 5, 5, 6},
	[]int{2, 3, 3, 2, 2, 2, 1, 1},
}

func check(nums []int) bool {
	size := len(nums)
	for i := 0; i < size; i = i + 2 {
		if i-1 > 0 && nums[i-1] <= nums[i] {
			return false
		}

		if i+1 < size && nums[i+1] <= nums[i] {
			return false
		}
	}

	return true
}

func Test_wiggleSort(t *testing.T) {
	ast := assert.New(t)
	for _, nums := range questions {
		temp := append([]int{}, nums...)
		wiggleSort(nums)
		ast.True(check(nums), "输入 %v, 输出 %v", temp, nums)
	}
}

func Benchmark_wiggleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, nums := range questions {
			wiggleSort(nums)
		}
	}
}
