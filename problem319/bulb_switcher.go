package problem319

// 参看 https://blog.csdn.net/baidu_23318869/article/details/50386323
func bulbSwitch(n int) int {
	res := 0
	for i := 1; i*i <= n; i++ {
		res++
	}

	return res
}
