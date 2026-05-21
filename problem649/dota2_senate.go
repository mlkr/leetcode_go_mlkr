package problem649

func predictPartyVictory(senate string) string {
	size := len(senate)
	if size == 1 {
		if senate == "D" {
			return "Dire"
		}

		return "Radiant"
	}

	bs := make([]byte, size)
	bs = []byte(senate)

	countD, countR := 0, 0
	i := 0
	for {
		if bs[i] == 'D' {
			countD++
			if countR > 0 {
				countR--
				bs[i] = 'B'
				countD--
			}
		}

		if bs[i] == 'R' {
			countR++
			if countD > 0 {
				countD--
				bs[i] = 'B'
				countR--
			}
		}

		i++
		if i == size {
			bs = getNewBs(bs)
			if size == len(bs) {
				// 没有可以ban的了， 全都是一个阵营
				break
			}

			size = len(bs)
			i = 0
		}
	}

	if bs[0] == 'D' {
		return "Dire"
	}

	return "Radiant"
}

// 去掉ban的
func getNewBs(bs []byte) []byte {
	size := len(bs)
	newBs := make([]byte, size)
	j := 0
	for i := 0; i < size; i++ {
		if bs[i] != 'B' {
			newBs[j] = bs[i]
			j++
		}
	}

	return newBs[:j]
}

// 这个最好
func predictPartyVictory2(senate string) string {
	size := len(senate)
	dque, rque := make([]int, 0, size), make([]int, 0, size)
	for i := 0; i < size; i++ {
		if senate[i] == 'D' {
			dque = append(dque, i)
		}

		if senate[i] == 'R' {
			rque = append(rque, i)
		}
	}

	for len(dque) > 0 && len(rque) > 0 {
		di := dque[0]
		dque = dque[1:]

		ri := rque[0]
		rque = rque[1:]

		if di < ri {
			dque = append(dque, size+di)
		}

		if ri < di {
			rque = append(rque, size+ri)
		}
	}

	if len(dque) == 0 {
		return "Radiant"
	}

	return "Dire"
}
