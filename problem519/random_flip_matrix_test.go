package problem519

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solution(t *testing.T) {
	ast := assert.New(t)
	obj := Constructor(3, 1)
	nums1 := obj.Flip()
	nums2 := obj.Flip()
	nums3 := obj.Flip()
	nums4 := obj.Flip()
	fmt.Printf("输出 %v %v %v %v\n", nums1, nums2, nums3, nums4)
	ast.Nil(nums4, "输出 %v", nums4)

	obj.Reset()
	nums1 = obj.Flip()
	nums2 = obj.Flip()
	nums3 = obj.Flip()
	nums4 = obj.Flip()
	fmt.Printf("输出 %v %v %v %v\n", nums1, nums2, nums3, nums4)

}

func Benchmark_Solution(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
