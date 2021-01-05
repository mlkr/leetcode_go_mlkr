package problem466

import "strings"

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	size1 := len(s1)
	size2 := len(s2)

	if size1*n1 < size2*n2 {
		return 0
	}

	s1 = strings.Repeat(s1, n1)
	s2 = strings.Repeat(s2, n2)

	j := 0
	res := 0
	for i := 0; i < size1*n1; i++ {
		if s1[i] == s2[j] {
			j++
			if j == size2*n2 {
				j = 0
				res++
			}
		}
	}

	return res
}

// 最佳
func getMaxRepetitions2(s1 string, n1 int, s2 string, n2 int) int {
	size1 := len(s1)
	size2 := len(s2)

	if size1*n1 < size2*n2 {
		return 0
	}

	last := make([]int, size2)
	count := make([]int, n1+1)

	j, cnt := 0, 0
	for k := 1; k <= n1; k++ {
		for i := 0; i < size1; i++ {
			if s1[i] == s2[j] {
				j++
				if j == size2 {
					j = 0
					cnt++
				}
			}
		}

		if last[j] == 0 {
			last[j] = k
			count[k] = cnt
			continue
		}

		// last[j] != 0 s2周期重复
		start := last[j]

		// 周期为 p
		p := k - start

		// s2 每个周期重复 t 遍
		t := cnt - count[start]

		// count[start+(n1-start)%p] 是剩下的，不能构成一个周期
		res := (n1-start)/p*t + count[start+(n1-start)%p]
		return res / n2
	}

	return cnt / n2
}
