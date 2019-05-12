package problem275

func hIndex(citations []int) int {
	size := len(citations)
	i, j := 0, size-1
	for i < j {
		citations[i], citations[j] = citations[j], citations[i]
		i++
		j--
	}

	// 二分查找
	lo, hi := 0, size-1
	for lo <= hi {
		mi := (lo + hi) / 2
		if citations[mi] > mi {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}

	return lo
}

// 解法二
func hIndex2(citations []int) int {
	// 二分查找
	size := len(citations)
	lo, hi := 0, size-1
	for lo <= hi {
		mi := (lo + hi) / 2
		if citations[mi] > size-1-mi {
			hi = mi - 1
		} else {
			lo = mi + 1
		}
	}

	return len(citations) - 1 - hi
}
