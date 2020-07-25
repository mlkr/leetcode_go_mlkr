package problem423

func originalDigits(s string) string {
	var count [10]int
	for _, b := range s {
		switch b {
		case 'z':
			count[0]++
		case 'w':
			count[2]++
		case 'u':
			count[4]++
		case 'x':
			count[6]++
		case 'g':
			count[8]++
		case 's': // six seven
			count[7]++
		case 'v': // five seven
			count[5]++
		case 'h': // three eight
			count[3]++
		case 'i': // five six eight nine
			count[9]++
		case 'o': // zero one two four
			count[1]++
		}
	}

	count[7] -= count[6]
	count[5] -= count[7]
	count[3] -= count[8]
	count[1] -= count[0] + count[2] + count[4]
	count[9] -= count[5] + count[6] + count[8]

	res := make([]byte, 0, len(s)/3)
	for i := byte(0); i < 10; i++ {
		for count[i] > 0 {
			res = append(res, i+'0')
			count[i]--
		}
	}

	return string(res)
}
