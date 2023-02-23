package dpTest

func MinPathSum1(m [][]int) int {
	if m == nil || len(m) == 0 || len(m[0]) == 0 || m[0] == nil {
		return 0
	}
	row, col := len(m), len(m[0])
	//新建dp矩阵 储存结果
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	dp[0][0] = m[0][0]
	//获取第一列的累加和并存入dp表中去
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + m[i][0]
	}
	//获取第一行的累加和
	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + m[0][j]
	}
	//普遍位置
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + m[i][j]
		}
	}
	return dp[row-1][col-1]
}
