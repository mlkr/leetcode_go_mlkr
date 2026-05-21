package problem649

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	senate, ans string
}{
	{
		"RRDRDDRDRRDDDDDRDRDR",
		"Radiant",
	},

	{
		"DDRRR",
		"Dire",
	},

	{
		"RD",
		"Radiant",
	},

	{
		"RDD",
		"Dire",
	},
}

func Test_predictPartyVictory(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, predictPartyVictory(q.senate), "输入 %v", q.senate)
	}
}

func Benchmark_predictPartyVictory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			predictPartyVictory(q.senate)
		}
	}
}

func Test_predictPartyVictory2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, predictPartyVictory2(q.senate), "输入 %v", q.senate)
	}
}

func Benchmark_predictPartyVictory2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			predictPartyVictory2(q.senate)
		}
	}
}
