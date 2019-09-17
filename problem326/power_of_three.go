package problem326

// 参看https://leetcode.com/problems/power-of-three/discuss/77856/1-line-java-solution-without-loop-recursion
func isPowerOfThree(n int) bool {
	// 1162261467 = 3^19, 3^20 比 32位机器的 int 型大
	return (n > 0 && 1162261467%n == 0)
}
