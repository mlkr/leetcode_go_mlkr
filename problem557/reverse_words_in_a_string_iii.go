package problem557

import "strings"

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	size := len(ss)
	res := make([]string, size)
	for i := 0; i < size; i++ {
		res[i] = reverseWord(ss[i])
	}

	return strings.Join(res, " ")
}

func reverseWord(word string) string {
	bs := []byte(word)
	i := 0
	j := len(bs) - 1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}

	return string(bs)
}
