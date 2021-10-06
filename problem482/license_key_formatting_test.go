package problem482

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	k   int
	ans string
}{
	{
		"---",
		3,
		"",
	},

	{
		"--a-a-a-a--",
		2,
		"AA-AA",
	},

	{
		"5F3Z-2e-9-w",
		4,
		"5F3Z-2E9W",
	},

	{
		"2-5g-3-J",
		2,
		"2-5G-3J",
	},
}

func Test_licenseKeyFormatting(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, licenseKeyFormatting(q.s, q.k), "输入 %v %v", q.s, q.k)
	}
}

func Benchmakr_licenseKeyFormatting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			licenseKeyFormatting(q.s, q.k)
		}
	}
}
