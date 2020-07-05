package problem406

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			// k 从小到大
			return people[i][1] < people[j][1]
		}

		// h 从高到低
		return people[i][0] > people[j][0]
	})

	res := make([][]int, 0, len(people))

	insert := func(idx int, person []int) {
		res = append(res, person)
		if idx == len(res)-1 {
			return
		}

		copy(res[idx+1:], res[idx:])
		res[idx] = person
	}

	for _, person := range people {
		insert(person[1], person)
	}

	return res
}

// 最优
func reconstructQueue2(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			// h 相同 k从多到少排
			return people[i][1] > people[j][1]
		}

		// h 从低到高排
		return people[i][0] < people[j][0]
	})

	indexes := make([]int, len(people))
	for i := range indexes {
		indexes[i] = i
	}

	queued := make([][]int, len(people))
	for _, person := range people {
		k := person[1]
		queued[indexes[k]] = person

		indexes = append(indexes[:k], indexes[k+1:]...)
	}

	return queued
}
