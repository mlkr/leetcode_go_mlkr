package problem443

import "strconv"

func compress(chars []byte) int {
	count := 1
	end := 0
	for i, char := range chars {
		if i+1 < len(chars) && char == chars[i+1] {
			count++
		} else {
			chars[end] = char
			end++

			if count > 1 {
				for _, num := range strconv.Itoa(count) {
					chars[end] = byte(num)
					end++
				}
			}

			count = 1
		}
	}

	return end
}
