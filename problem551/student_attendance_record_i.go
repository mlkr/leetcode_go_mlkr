package problem551

func checkRecord(s string) bool {
	Acount := 0
	Lcount := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			Acount++
			if Acount >= 2 {
				return false
			}
		}

		if s[i] == 'L' {
			Lcount++
			if Lcount == 3 {
				return false
			}

			continue
		}

		Lcount = 0
	}

	return true
}
