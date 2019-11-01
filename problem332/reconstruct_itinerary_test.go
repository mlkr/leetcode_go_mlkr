package problem332

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	tickets [][]string
	ans     []string
}{
	{
		[][]string{
			[]string{"MUC", "LHR"},
			[]string{"JFK", "MUC"},
			[]string{"SFO", "SJC"},
			[]string{"LHR", "SFO"},
		},

		[]string{"JFK", "MUC", "LHR", "SFO", "SJC"},
	},

	{
		[][]string{
			[]string{"JFK", "SFO"},
			[]string{"JFK", "ATL"},
			[]string{"SFO", "ATL"},
			[]string{"ATL", "JFK"},
			[]string{"ATL", "SFO"},
		},

		[]string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
	},

	{
		[][]string{
			[]string{"EZE", "TIA"},
			[]string{"EZE", "HBA"},
			[]string{"AXA", "TIA"},
			[]string{"JFK", "AXA"},
			[]string{"ANU", "JFK"},
			[]string{"ADL", "ANU"},
			[]string{"TIA", "AUA"},
			[]string{"ANU", "AUA"},
			[]string{"ADL", "EZE"},
			[]string{"ADL", "EZE"},
			[]string{"EZE", "ADL"},
			[]string{"AXA", "EZE"},
			[]string{"AUA", "AXA"},
			[]string{"JFK", "AXA"},
			[]string{"AXA", "AUA"},
			[]string{"AUA", "ADL"},
			[]string{"ANU", "EZE"},
			[]string{"TIA", "ADL"},
			[]string{"EZE", "ANU"},
			[]string{"AUA", "ANU"},
		},

		[]string{"JFK", "AXA", "AUA", "ADL", "ANU", "AUA", "ANU", "EZE", "ADL", "EZE", "ANU", "JFK", "AXA", "EZE", "TIA", "AUA", "AXA", "TIA", "ADL", "EZE", "HBA"},
	},
}

func Test_findItinerary(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findItinerary(q.tickets), "expected: %v", q.ans)
	}
}

func Benchmark_findItinerary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findItinerary(q.tickets)
		}
	}
}
