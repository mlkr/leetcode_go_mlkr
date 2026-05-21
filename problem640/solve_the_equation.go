package problem640

import (
	"strconv"
	"strings"
)

func solveEquation(equation string) string {
	ss := strings.Split(equation, "=")
	la, lb := analyze(ss[0])
	ra, rb := analyze(ss[1])

	a := la - ra
	b := rb - lb

	if a == 0 {
		if b == 0 {
			return "Infinite solutions"
		}
		return "No solution"
	}

	return "x=" + strconv.Itoa(b/a)
}

func analyze(equ string) (int, int) {
	equ = strings.Replace(equ, "-", "+-", -1)
	ss := strings.Split(equ, "+")

	a := 0
	b := 0
	for _, s := range ss {
		if s == "" {
			continue
		}

		if s == "x" {
			a++
			continue
		}

		if s == "-x" {
			a--
			continue
		}

		if s[len(s)-1] == 'x' {
			num, _ := strconv.Atoi(s[:len(s)-1])
			a += num
			continue
		}

		num, _ := strconv.Atoi(s)
		b += num
	}

	return a, b
}
