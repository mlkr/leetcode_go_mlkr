package problem165

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para struct {
		ver1 string
		ver2 string
	}
	ans int
}{
	{
		struct {
			ver1 string
			ver2 string
		}{
			"1", "1.1",
		},
		-1,
	},

	{
		struct {
			ver1 string
			ver2 string
		}{
			"1.0", "1",
		},
		0,
	},

	{
		struct {
			ver1 string
			ver2 string
		}{
			"01", "1",
		},
		0,
	},

	{
		struct {
			ver1 string
			ver2 string
		}{
			"0.1", "1.1",
		},
		-1,
	},

	{
		struct {
			ver1 string
			ver2 string
		}{
			"1.0.1", "1",
		},
		1,
	},

	{
		struct {
			ver1 string
			ver2 string
		}{
			"7.5.2.4", "7.5.3",
		},
		-1,
	},
}

func Test_compareVersion(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(compareVersion(q.para.ver1, q.para.ver2), q.ans)
	}
}

func Benchmark_compareVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			compareVersion(q.para.ver1, q.para.ver2)
		}
	}
}
