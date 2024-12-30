package problem609

import "strings"

func findDuplicate(paths []string) [][]string {
	m := make(map[string][]string)

	for _, path := range paths {
		fs := strings.Split(path, " ")
		folder := fs[0]
		fs = fs[1:]

		for _, file := range fs {
			i := strings.IndexByte(file, '(')
			name := file[:i]
			content := file[i+1 : len(file)-1]

			m[content] = append(m[content], folder+"/"+name)
		}
	}

	res := make([][]string, 0, len(m))
	for _, files := range m {
		if len(files) > 1 {
			res = append(res, files)
		}
	}

	return res
}
