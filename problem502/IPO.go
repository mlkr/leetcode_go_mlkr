package problem502

import "container/heap"

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	availableProjects := &Project{
		weight:  profits,
		reverse: true,
	}

	pendingProjects := &Project{
		weight: capital,
	}

	for i, v := range capital {
		if w < v {
			heap.Push(pendingProjects, i)
			continue
		}
		heap.Push(availableProjects, i)
	}

	for k > 0 {
		k--

		if availableProjects.Len() == 0 {
			break
		}

		idx := heap.Pop(availableProjects).(int)
		w += profits[idx]

		for pendingProjects.Len() > 0 && w >= capital[pendingProjects.indexes[0]] {
			heap.Push(availableProjects, heap.Pop(pendingProjects).(int))
		}
	}

	return w
}

type Project struct {
	weight  []int
	indexes []int
	reverse bool
}

func (p Project) Len() int { return len(p.indexes) }

func (p Project) Swap(i, j int) { p.indexes[i], p.indexes[j] = p.indexes[j], p.indexes[i] }

func (p Project) Less(i, j int) bool {
	if p.reverse {
		return p.weight[p.indexes[i]] > p.weight[p.indexes[j]]
	}
	return p.weight[p.indexes[i]] < p.weight[p.indexes[j]]
}

func (p *Project) Push(x interface{}) {
	index := x.(int)
	p.indexes = append(p.indexes, index)
}

func (p *Project) Pop() interface{} {
	size := len(p.indexes)
	item := p.indexes[size-1]
	p.indexes = p.indexes[:size-1]

	return item
}
