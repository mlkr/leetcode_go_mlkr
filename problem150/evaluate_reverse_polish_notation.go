package problem150

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stock := make([]int, 0, len(tokens))

	for _, s := range tokens {
		num, err := strconv.Atoi(s)
		if err == nil {
			stock = append(stock, num)
		} else {
			switch s {
			case "+":
				stock[len(stock)-2] += stock[len(stock)-1]
			case "-":
				stock[len(stock)-2] -= stock[len(stock)-1]
			case "*":
				stock[len(stock)-2] *= stock[len(stock)-1]
			case "/":
				stock[len(stock)-2] /= stock[len(stock)-1]
			}

			stock = stock[:len(stock)-1]
		}
	}

	return stock[0]
}
