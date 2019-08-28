package problem316

import "strings"

func removeDuplicateLetters(s string) string {
	if len(s) == 0 {
		return ""
	}

	counts := [26]int{}
	for i := range s {
		counts[s[i]-'a']++
	}

	begin := 0
	for i := range s {
		if s[i] < s[begin] {
			begin = i
		}

		counts[s[i]-'a']--
		if counts[s[i]-'a'] == 0 {
			break
		}
	}

	return string(s[begin]) + removeDuplicateLetters(strings.Replace(s[begin+1:], string(s[begin]), "", -1))
}

func removeDuplicateLetters2(s string) string {
	counts := [26]int{}
	for i := range s {
		counts[s[i]-'a']++
	}

	visited := [26]bool{}
	res := make([]byte, 0, len(s))
	for i := range s {
		counts[s[i]-'a']--

		if visited[s[i]-'a'] {
			continue
		}

		for len(res) > 0 && res[len(res)-1] > s[i] && counts[res[len(res)-1]-'a'] > 0 {
			visited[res[len(res)-1]-'a'] = false
			res = res[:len(res)-1]
		}

		res = append(res, s[i])
		visited[res[len(res)-1]-'a'] = true
	}

	return string(res)
}
