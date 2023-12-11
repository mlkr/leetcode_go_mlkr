package problem541

func reverseStr(s string, k int) string {
	bs := []byte(s)
	size := len(bs)

	i := 0
	for {
		begin := i * k
		if begin >= size {
			break
		}

		end := i*k + k
		if end > size {
			end = size
		}

		rever(bs[begin:end])

		i += 2
	}

	return string(bs)
}

func rever(bs []byte) {
	size := len(bs)
	i := 0
	j := size - 1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
}
