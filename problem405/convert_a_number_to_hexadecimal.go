package problem405

func toHex(num int) string {
	bs := []byte{'0','1','2','3','4','5','6','7','8','9','a','b','c','d','e','f'}

	if num == 0 {
		return "0"
	}

	if num < 0 {
		// 2的32次方 = 4294967296
		num += 4294967296
	}

	res := []byte{}
	for num > 0{
		idx:=num%16
		res = append([]byte{bs[idx]}, res...)

		num/=16
	}

	return string(res)
}

func toHex2(num int) string {
	if num == 0 {
		return "0"
	}

	res := ""
	h := "0123456789abcdef"
	for i:=0; num!=0 && i<8;i++{
		res = string(h[num&15]) + res
		num >>= 4
	}

	return res
}