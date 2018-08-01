package problem67

func addBinary(a string, b string) string {
	aInt := sConvInts(a)
	bInt := sConvInts(b)

	aLen := len(aInt)
	bLen := len(bInt)

	if aLen == 0 {
		return b
	}

	if bLen == 0 {
		return a
	}

	if aLen > bLen {
		bInt = append(make([]int, aLen-bLen), bInt...)
	}

	if bLen > aLen {
		aInt = append(make([]int, bLen-aLen), aInt...)
	}

	length := max(aLen, bLen)
	for j := length - 1; j > 0; j-- {
		aInt[j] += bInt[j]
		if aInt[j] > 1 {
			aInt[j-1]++
			aInt[j] = aInt[j] - 2
		}
	}

	aInt[0] += bInt[0]
	if aInt[0] > 1 {
		aInt[0] = aInt[0] - 2
		aInt = append([]int{1}, aInt...)
	}

	return intsConvS(aInt)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 字符串转数组
func sConvInts(s string) []int {
	res := make([]int, len(s))
	for i := range s {
		temp := 0
		if s[i] == '1' {
			temp = 1
		}

		res[i] = temp
	}

	return res
}

// 数组转字符串
func intsConvS(arr []int) string {
	res := make([]byte, len(arr))
	for i := range arr {
		temp := '0'
		if arr[i] == 1 {
			temp = '1'
		}

		res[i] = byte(temp)
	}

	return string(res)
}
