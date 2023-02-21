package dpTest

//机器人到达指定位置的最大方法数
//位置：1~N 机器人位于M M∈[1,N]
//  机器人可以往左走往右走
//规定 走K 步 到达 p 返回方法数

// N 固定参数
// cur 当前位置
// rest 剩余步数
// p 终止位置 固定参数
func walk(N, cur, rest, p int) int {
	// 如果剩余的步数为0 cur若等于 p 方法数++
	if rest == 0 {
		if cur == p {
			return 1
		}
		return 0
	}
	//当cur来到了1位置
	if cur == 1 {
		//只能往右走
		return walk(N, 2, rest-1, p)
	}
	//cur 来到了N位置
	if cur == N {
		//只能往左走
		return walk(N, N-1, rest-1, p)
	}

	return walk(N, cur-1, rest-1, p) + walk(N, cur+1, rest-1, p)

}
func Way1(N, M, K, P int) int {
	if N < 2 || K < 1 || M < 1 || M > N || P < 1 || P > N {
		return 0
	}
	return walk(N, M, K, P)
}
func Way2(N, M, K, P int) int {
	if N < 2 || K < 1 || M < 1 || M > N || P < 1 || P > N {
		return 0
	}
	dp := make([][]int, K+1)
	for i, _ := range dp {
		dp[i] = make([]int, N+1)
	}
	dp[0][P] = 1
	for i := 1; i <= K; i++ {
		for j := 1; j <= N; j++ {
			if j == 1 {
				dp[i][j] = dp[i-1][2]
			} else if j == N {
				dp[i][j] = dp[i-1][N-1]
			} else {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j+1]
			}
		}
	}
	return dp[K][M]
}
func Way3(N, M, K, P int) int {
	if N < 2 || K < 1 || M < 1 || M > N || P < 1 || P > N {
		return 0
	}
	dp := make([][]int, K+1)

	for i := 0; i <= K; i++ {
		dp[i] = make([]int, N+1)
		for j := 0; j <= N; j++ {
			dp[i][j] = -1
		}
	}
	return walk2(dp, N, M, K, P)
}
func walk2(dp [][]int, N, cur, rest, p int) int {
	if dp[rest][cur] != -1 {
		return dp[rest][cur]
	}
	// 如果剩余的步数为0 cur若等于 p 方法数++
	if rest == 0 {
		if cur == p {
			dp[rest][cur] = 1
			return dp[rest][cur]
		}
		dp[rest][cur] = 0
		return dp[rest][cur]
	}
	//当cur来到了1位置
	if cur == 1 {
		//只能往右走
		dp[rest][cur] = walk2(dp, N, 2, rest-1, p)
		return dp[rest][cur]
	}
	//cur 来到了N位置
	if cur == N {
		//只能往左走
		dp[rest][cur] = walk2(dp, N, N-1, rest-1, p)
		return dp[rest][cur]
	}
	dp[rest][cur] = walk2(dp, N, cur-1, rest-1, p) + walk2(dp, N, cur+1, rest-1, p)
	return dp[rest][cur]
}
