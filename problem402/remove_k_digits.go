package problem402

func removeKdigits(num string, k int) string {
	stack := make([]byte, len(num))
	top := 0
	right := len(num) - k
	if right == 0 {
		return "0"
	}

	for i:=range num {
		for top > 0 && stack[top-1] > num[i] && k > 0 {
			top--
			k--
		}

		stack[top] = num[i]
		top++
	}

	left:=0
	for stack[left] == '0' {
		left++

		if left == right {
			return "0"
		}
	}


	return string(stack[left:right])
}
