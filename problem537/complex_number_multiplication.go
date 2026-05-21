package problem537

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	a, b := getRealImg(num1)
	c, d := getRealImg(num2)

	return fmt.Sprintf("%d+%di", a*c-b*d, b*c+a*d)
}

func getRealImg(num string) (real, img int) {
	ss := strings.Split(num, "+")

	real, _ = strconv.Atoi(ss[0])
	img, _ = strconv.Atoi(ss[1][:len(ss[1])-1])

	return
}
