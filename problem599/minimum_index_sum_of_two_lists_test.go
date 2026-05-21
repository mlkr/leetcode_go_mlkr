package problem599

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	list1, list2, ans []string
}{
	{
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
		[]string{"Shogun"},
	},

	{
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"KFC", "Shogun", "Burger King"},
		[]string{"Shogun"},
	},

	{
		[]string{"happy", "sad", "good"},
		[]string{"sad", "happy", "good"},
		[]string{"sad", "happy"},
	},
}

func Test_findRestaurant(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.ElementsMatch(q.ans, findRestaurant(q.list1, q.list2))
	}
}

func Benchmark_findRestaurant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findRestaurant(q.list1, q.list2)
		}
	}
}
