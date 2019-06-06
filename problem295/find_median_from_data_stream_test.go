package problem295

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MedianFinder(t *testing.T) {
	ast := assert.New(t)

	m := Constructor()
	m.AddNum(1)
	m.AddNum(2)
	ast.Equal(1.5, m.FindMedian())

	m.AddNum(3)
	ast.Equal(2.0, m.FindMedian())

	m.AddNum(35)
	m.AddNum(30)

	ast.Equal(3.0, m.FindMedian())

}

func Test_MedianFinder2(t *testing.T) {
	ast := assert.New(t)

	m := Constructor2()
	m.AddNum2(1)
	m.AddNum2(2)
	ast.Equal(1.5, m.FindMedian2())

	m.AddNum2(3)
	ast.Equal(2.0, m.FindMedian2())

	m.AddNum2(35)
	m.AddNum2(30)

	ast.Equal(3.0, m.FindMedian2())

	// 为了使代码覆盖率 100%
	m.left.Pop()

	m1 := Constructor2()
	m1.AddNum2(-1)
	ast.Equal(-1.0, m1.FindMedian2())

	m1.AddNum2(-2)
	ast.Equal(-1.5, m1.FindMedian2())

	m1.AddNum2(-3)
	ast.Equal(-2.0, m1.FindMedian2())

	m1.AddNum2(-4)
	ast.Equal(-2.5, m1.FindMedian2())

	m1.AddNum2(-5)
	ast.Equal(-3.0, m1.FindMedian2())

}

func Benchmark_MedianFinder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := Constructor()
		m.AddNum(1)
		m.AddNum(2)
		m.FindMedian()

		m.AddNum(3)
		m.FindMedian()
	}
}

func Benchmark_MedianFinder2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := Constructor2()
		m.AddNum2(1)
		m.AddNum2(2)
		m.FindMedian2()

		m.AddNum2(3)
		m.FindMedian2()
	}
}
