package problem332

import "sort"

func findItinerary(tickets [][]string) []string {
	nexts := make(map[string][]string)
	for _, ticket := range tickets {
		nexts[ticket[0]] = append(nexts[ticket[0]], ticket[1])
	}

	for _, next := range nexts {
		sort.Strings(next)
	}

	path := make([]string, 0, len(tickets)+1)

	var visit func(airport string)
	visit = func(airport string) {
		for len(nexts[airport]) > 0 {
			next := nexts[airport][0]
			nexts[airport] = nexts[airport][1:]

			visit(next)

		}

		path = append(path, airport)
	}

	visit("JFK")

	i, j := 0, len(path)-1
	for i < j {
		path[i], path[j] = path[j], path[i]
		i++
		j--
	}

	return path
}
