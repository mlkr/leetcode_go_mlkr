package problem564

import (
	"strconv"
)

func nearestPalindromic(n string) string {
	size := len(n)
	candidates := append(getSizeCandidates(size), getCandidates(n)...)

	num, _ := strconv.Atoi(n)
	delta := func(x int) int {
		if num > x {
			return num - x
		}

		return x - num
	}

	res := 1<<63 - 1
	for _, candidate := range candidates {
		if num == candidate {
			continue
		}

		if delta(res) > delta(candidate) ||
			delta(res) == delta(candidate) && res > candidate {
			res = candidate
		}
	}

	return strconv.Itoa(res)
}

func getCandidates(n string) []int {
	size := len(n)
	prefix, _ := strconv.Atoi(n[:(size+1)/2])

	res, ps := make([]int, 3), make([]int, 3)
	res[0], ps[0] = prefix-1, prefix-1
	res[1], ps[1] = prefix, prefix
	res[2], ps[2] = prefix+1, prefix+1

	if size%2 == 1 {
		for i := range ps {
			ps[i] /= 10
		}
	}

	for i := range res {
		for ps[i] > 0 {
			res[i] = res[i]*10 + ps[i]%10
			ps[i] /= 10
		}
	}

	return res
}

func getSizeCandidates(size int) []int {
	base := 1
	for i := 1; i < size; i++ {
		base *= 10
	}

	res := make([]int, 4)
	res[0] = base - 1
	res[1] = base + 1

	base *= 10

	res[2] = base - 1
	res[3] = base + 1

	return res
}
