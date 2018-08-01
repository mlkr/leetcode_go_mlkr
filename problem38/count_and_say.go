package problem38

func countAndSay(n int) string {
	buf := []byte{'1'}

	for n > 1 {
		buf = say(buf)
		n--
	}

	return string(buf)
}

func say(buf []byte) []byte {
	res := make([]byte, 0, len(buf)*2)

	i, j := 0, 1

	for i < len(buf) {
		// 计算数量
		for j < len(buf) && buf[i] == buf[j] {
			j++
		}

		res = append(res, byte(j-i+'0'), buf[i])

		i = j
	}

	return res
}
