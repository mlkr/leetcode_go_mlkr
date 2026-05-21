package problem591

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	code string
	ans  bool
}{
	{
		"<123456>",
		false,
	},

	{
		"<DIV></DIV",
		false,
	},

	{
		"12",
		false,
	},

	{
		"<12345678",
		false,
	},

	{
		"<1234567890></1234567890>",
		false,
	},

	{

		"</A></A></A></A>",
		false,
	},

	{
		"<DIV>This is the first line <![CDATA[<div>]]></DIV>",
		true,
	},

	{
		"<DIV>This is the first line <![CDATA[<div>]</DIV>",
		false,
	},

	{
		"<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>",
		true,
	},

	{
		"<A>  <B> </A>   </B>",
		false,
	},

	{
		"<DIV>  div tag is not closed  <DIV>",
		false,
	},

	{
		"<DIV>  unmatched <  </DIV>",
		false,
	},

	{
		"<DIV> closed tags with invalid tag name  <b>123</b> </DIV>",
		false,
	},

	{
		"<DIV> unmatched tags with invalid tag name  </1234567890> and <CDATA[[]]>  </DIV>",
		false,
	},

	{
		"<DIV>  unmatched start tag <B>  and unmatched end tag </C>  </DIV>",
		false,
	},
}

func Test_isValid(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, isValid(q.code), "输入 %v", q.code)
	}
}

func Benchmark_isValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isValid(q.code)
		}
	}
}
