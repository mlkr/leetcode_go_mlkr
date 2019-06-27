// 参看
// https://blog.csdn.net/fuxuemingzhu/article/details/83537572

package problem304

type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var dp [][]int
	m, n := 0, 0
	m = len(matrix)
	if m > 0 {
		n = len(matrix[0])
	}

	if m == 0 && n == 0 {
		return NumMatrix{
			dp: nil,
		}
	}

	dp = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = matrix[i-1][j-1] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
		}
	}

	return NumMatrix{
		dp: dp,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if this.dp == nil {
		return 0
	}
	return this.dp[row2+1][col2+1] - this.dp[row2+1][col1] - this.dp[row1][col2+1] + this.dp[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
