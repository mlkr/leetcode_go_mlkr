package problem468

import (
	"strconv"
	"strings"
)

var ans = []string{"IPv4", "IPv6", "Neither"}

func validIPAddress(IP string) string {
	// IPv4
	IPArr := strings.Split(IP, ".")
	size := len(IPArr)

	if size == 4 {
		for i := 0; i < size; i++ {
			str := IPArr[i]

			// 范围在 0~255 之间
			num, err := strconv.Atoi(str)
			if err != nil || (num < 0 || num > 255) {
				return ans[2]
			}

			// 不能有前导0
			if str[0] == '0' && len(str) > 1 {
				return ans[2]
			}
		}

		return ans[0]
	}

	// IPv6
	IPArr = strings.Split(IP, ":")
	size = len(IPArr)

	if size == 8 {
		for i := 0; i < size; i++ {
			str := IPArr[i]

			// 长度在 1~4 之间
			if len(str) < 1 || len(str) > 4 {
				return ans[2]
			}

			// 解析为16进制数
			_, err := strconv.ParseInt(str, 16, 0)
			if err != nil {
				return ans[2]
			}
		}

		return ans[1]

	}

	return ans[2]
}
