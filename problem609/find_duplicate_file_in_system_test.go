package problem609

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	paths []string
	ans   [][]string
}{
	{
		[]string{"root/a 1.txt(abcd) 2.txt(efsfgh)", "root/c 3.txt(abdfcd)", "root/c/d 4.txt(efggdfh)"},
		[][]string{},
	},

	{
		[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"},
		[][]string{
			[]string{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"},
			[]string{"root/a/1.txt", "root/c/3.txt"},
		},
	},

	{
		[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)"},
		[][]string{
			[]string{"root/a/2.txt", "root/c/d/4.txt"},
			[]string{"root/a/1.txt", "root/c/3.txt"},
		},
	},
}

func Test_findDuplicate(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.ElementsMatch(q.ans, findDuplicate(q.paths), "输入 %v", q.paths)
	}
}

func Benchmark_findDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findDuplicate(q.paths)
		}
	}
}

// 打孔机
// 外接光驱
// 国产平板
