package problem591

import "strings"

func isValid(code string) bool {

	var validTagName func(name string) bool
	validTagName = func(name string) bool {
		size := len(name)
		if size < 1 || size > 9 {
			return false
		}

		for i := 0; i < size; i++ {
			if name[i] < 'A' || name[i] > 'Z' {
				return false
			}
		}

		return true
	}

	stack := []string{}

	for i := 0; i < len(code); i++ {
		if i > 0 && len(stack) == 0 {
			return false
		}

		switch {
		case strings.HasPrefix(code[i:], "<![CDATA["):
			j := i + strings.Index(code[i:], "]]>")
			if j < i {
				return false
			}
			i = j + 2

		case strings.HasPrefix(code[i:], "</"):
			i += strings.Index(code[i:], "</") + 2
			j := i + strings.Index(code[i:], ">")
			if j < i {
				return false
			}

			if len(stack) == 0 || stack[len(stack)-1] != code[i:j] {
				return false
			}

			stack = stack[:len(stack)-1]
			i = j

		case code[i] == '<':
			i++
			j := i + strings.Index(code[i:], ">")
			if j < i {
				return false
			}

			if !validTagName(code[i:j]) {
				return false
			}

			stack = append(stack, code[i:j])
			i = j
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
