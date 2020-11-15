package problem458

// 参看 https://blog.csdn.net/fuxuemingzhu/article/details/81079261
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	// 每头猪可以测试的份数
	tests := minutesToTest/minutesToDie + 1

	res := 0
	nums := 1
	for buckets > nums {
		nums *= tests
		res++
	}

	return res
}
