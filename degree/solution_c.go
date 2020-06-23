package degree

import "math"

func InnerCalculateC(floor, cup int) (degree int) {
	dp := make([][]int, cup+1)
	for i := 0; i < cup+1; i++ {
		dp[i] = make([]int, floor+1)
	}
	for i := 1; i < floor+1; i++ {
		dp[1][i] = i
	}
	for i := 1; i < cup+1; i++ {
		dp[i][1] = 1
	}

	for i := 2; i < cup+1; i++ {
		for j := 2; j < floor+1; j++ {
			res := math.MaxInt64
			for k := 1; k < j+1; k++ {
				tmp := max(dp[i-1][k-1], dp[i][j-k])
				res = min(tmp, res)
			}
			dp[i][j] = res + 1
		}
	}

	return dp[cup][floor]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
