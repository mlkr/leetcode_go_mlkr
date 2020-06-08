package problem389

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var questions = []struct{
	s, t string
	ans byte
}{
	{
		"abcd",
		"abcde",
		'e',
	},
}

func Test_findTheDifference(t *testing.T){
	ast := assert.New(t)

	for _, q := range questions{
		ast.Equal(q.ans, findTheDifference(q.s, q.t), "输入 %s", q.s)
	}
}

func Benchmark_findTheDifference(b *testing.B) {
	for i:=0;i<b.N;i++{
		for _, q := range questions{
			findTheDifference(q.s, q.t)
		}		
	}
}