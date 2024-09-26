package problem592

import (
	"strconv"
	"strings"
)

func fractionAddition(expression string) string {
	exp := insertPlus(expression)
	ss := strings.Split(exp, "+")
	res := getFrac(ss[0])
	for i := 1; i < len(ss); i++ {
		res = add(res, getFrac(ss[i]))
	}

	return res.String()
}

type Frac struct {
	num int
	den int
}

func (f Frac) String() string {
	return strconv.Itoa(f.num) + "/" + strconv.Itoa(f.den)
}

func getFrac(exp string) Frac {
	ss := strings.Split(exp, "/")
	num, _ := strconv.Atoi(ss[0])
	den, _ := strconv.Atoi(ss[1])

	return Frac{
		num: num,
		den: den,
	}
}

func add(f1, f2 Frac) Frac {
	num := f1.num*f2.den + f2.num*f1.den
	den := f1.den * f2.den
	n := gcd(abs(num), den)

	return Frac{
		num: num / n,
		den: den / n,
	}
}

// a, b 最大公约数
//  12 8
//  8  4
//  4  0
func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}

	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

// 在expression[1:]插入+，将减法运算都转变为加法
func insertPlus(expression string) string {
	exp := []byte(expression)
	size := len(exp)
	res := make([]byte, 1, size)

	res[0] = exp[0]
	for i := 1; i < size; i++ {
		if exp[i] == '-' {
			res = append(res, '+')
		}

		res = append(res, exp[i])
	}

	return string(res)
}
