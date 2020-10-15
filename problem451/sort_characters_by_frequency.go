package problem451

import (
	"sort"
	"strings"
)

func frequencySort(s string) string {
	freq := ['z' + 1]int{}
	for _, b := range s {
		freq[b]++
	}

	ss := make([]string, 0, len(s))
	for k, v := range freq {
		if v == 0 {
			continue
		}

		ss = append(ss, strings.Repeat(string(k), v))
	}

	sort.Slice(ss, func(i, j int) bool {
		return len(ss[i]) > len(ss[j])
	})

	res := ""
	for _, s := range ss {
		res += s
	}

	return res
}

func frequencySort2(s string) string {
	freq := ['z' + 1]int{}
	for _, b := range s {
		freq[b]++
	}

	ss := make([]string, len(s)+1)
	for k, v := range freq {
		ss[v] += strings.Repeat(string(k), v)
	}

	sb := strings.Builder{}
	for i := len(s); i >= 0; i-- {
		sb.WriteString(ss[i])
	}

	return sb.String()
}
