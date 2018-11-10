package problem147

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para []int
	ans  []int
}{
	{
		[]int{},
		[]int{},
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

func Test_insertionSortList(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		l := sliceToList(q.para)
		sortList := insertionSortList(l)
		ast.Equal(listToSlice(sortList), q.ans)
	}
}

func Test_insertionSortList2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		l := sliceToList(q.para)
		sortList := insertionSortList2(l)
		ast.Equal(listToSlice(sortList), q.ans)
	}
}

func Benchmark_insertionSortList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			l := sliceToList(q.para)
			sortList := insertionSortList(l)
			listToSlice(sortList)
		}
	}
}

func Benchmark_insertionSortList2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			l := sliceToList(q.para)
			sortList := insertionSortList2(l)
			listToSlice(sortList)
		}
	}
}

func sliceToList(s []int) *ListNode {
	sLen := len(s)
	if sLen == 0 {
		return nil
	}

	cur := &ListNode{
		Val: s[0],
	}
	res := cur
	for i := 1; i < sLen; i++ {
		cur.Next = &ListNode{
			Val: s[i],
		}

		cur = cur.Next
	}

	return res
}

func listToSlice(l *ListNode) []int {
	res := []int{}
	for {
		if l == nil {
			break
		}

		res = append(res, l.Val)
		l = l.Next
	}

	return res
}
