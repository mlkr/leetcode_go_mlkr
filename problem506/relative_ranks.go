package problem506

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	p := Project{
		score: score,
	}

	size := len(score)
	p.indexes = make([]int, size)
	for i := 0; i < size; i++ {
		p.indexes[i] = i
	}

	sort.Sort(p)

	res := make([]string, size)
	for i := 0; i < size; i++ {
		if i == 0 {
			res[p.indexes[i]] = "Gold Medal"
			continue
		}

		if i == 1 {
			res[p.indexes[i]] = "Silver Medal"
			continue
		}

		if i == 2 {
			res[p.indexes[i]] = "Bronze Medal"
			continue
		}

		res[p.indexes[i]] = strconv.Itoa(i + 1)
	}

	return res
}

type Project struct {
	score   []int
	indexes []int
}

func (p Project) Len() int { return len(p.indexes) }

func (p Project) Swap(i, j int) { p.indexes[i], p.indexes[j] = p.indexes[j], p.indexes[i] }

func (p Project) Less(i, j int) bool {
	return p.score[p.indexes[i]] > p.score[p.indexes[j]]
}
