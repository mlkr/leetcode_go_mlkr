package problem71

import (
	"strings"
)

func simplifyPath(path string) string {
	pathLen := len(path)
	dir := make([]byte, 0, pathLen)
	stack := make([]string, 0, pathLen/2)

	for i := 0; i < pathLen; i++ {
		dir = dir[:0]

		for path[i] != '/' {
			dir = append(dir, path[i])
			i++

			if i == pathLen {
				break
			}
		}

		s := string(dir)

		switch s {
		case "", ".":
			// 什么也不做
		case "..":
			if len(stack) > 0 {
				// 回退上一级目录
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, s)
		}
	}

	return "/" + strings.Join(stack, "/")
}
