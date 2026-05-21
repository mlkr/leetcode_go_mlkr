package problem541

func reverseStr(s string, k int) string {
	bs := []byte(s)
	size := len(s)

	i := 0
	for {
		begin := i
		end := i + k
		if end > size {
			end = size
		}

		rever(bs[begin:end])
		i += 2 * k

		if i >= size {
			break
		}
	}

	return string(bs)
}

func rever(bs []byte) {
	i, j := 0, len(bs)-1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
}
