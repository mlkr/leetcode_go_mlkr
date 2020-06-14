package problem393

func validUtf8(data []int) bool {
	if len(data) == 0 {
		return false
	}

	// 剩下要检查的 byte 数
	cnt := 0

	for _, d := range data {
		if cnt == 0 {
			switch {
			case d>>3 == 30:
				cnt = 3
			case d>>4 == 14:
				cnt = 2
			case d>>5 == 6:
				cnt = 1
			case d>>7 == 0:
				continue
			default:
				return false
			}

		} else {
			if d>>6 != 2 {
				return false
			}

			cnt--
		}
	}

	return cnt == 0
}
