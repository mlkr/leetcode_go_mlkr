package problem415

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct{
	num1, num2, ans string
}{
	{
		"9",
		"99",
		"108",
	},

	{
		"6",
		"4",
		"10",
	},

	{
		"6",
		"40",
		"46",
	},

	{
		"3876620623801494171",
		"6529364523802684779",
		"10405985147604178950",
	},

	{
		"119",
		"10",
		"129",
	},
}

func Test_addStrings(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, addStrings(q.num1,q.num2), "输入 %v %v", q.num1,q.num2)
	}
}

func Benchmark_addStrings(b *testing.B) {
	for i:=0;i<b.N;i++{
		for _, q := range questions {
			addStrings(q.num1,q.num2)
		}
	}
}