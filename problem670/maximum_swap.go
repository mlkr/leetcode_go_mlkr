package problem670

func maximumSwap(num int) int {
	if num < 10 {
		return num
	}

	ns := num2ns(num)

	// indexes 包含所有数字最后出现的位置
	size := len(ns)
	indexes := make(map[int]int, size)
	for i := range ns {
		indexes[ns[i]] = i
	}

	// 将大的数字往前移
	for i := 0; i < size; i++ {
		for max := 9; max > ns[i]; max-- {
			j := indexes[max]
			if j > i {
				ns[i], ns[j] = ns[j], ns[i]
				return ns2num(ns)
			}
		}
	}

	return ns2num(ns)
}

func num2ns(num int) []int {
	ns := []int{}
	for num > 0 {
		ns = append(ns, num%10)
		num /= 10
	}

	i, j := 0, len(ns)-1
	for i < j {
		ns[i], ns[j] = ns[j], ns[i]
		i++
		j--
	}

	return ns
}

func ns2num(ns []int) int {
	num := 0
	for i := range ns {
		num = num*10 + ns[i]
	}

	return num
}
