package main

import (
	"fmt"
	"math"
)

// mathパッケージのMaxはfloat型しかないためint用にラップしておく
func max(lhs, rhs int) int {
	return int(math.Max(float64(lhs), float64(rhs)))
}

func main() {
	var (
		N, M int
	)
	// 入力
	fmt.Scanf("%d %d", &N, &M)
	values, weights := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d", &weights[i], &values[i])
	}
	// 32bitだとオーバーフローする場合があるので64bit
	// Goはデフォルトで0で初期化される.
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, M+1)
	}
	// 動的計画法部分
	// 緩和式にしたがって値を更新していく
	for i := 1; i <= N; i++ {
		for j := int(0); j <= M; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= weights[i-1] {
				fmt.Printf("dp[%d][%d]: %d\n", i, j, dp[i][j])
				fmt.Printf("dp[%d-1][%d-weights[%d-1]]: %d\n", i, j, i, dp[i-1][j-weights[i-1]])
				fmt.Printf("values[%d-1]: %d\n", i, values[i-1])
				dp[i][j] = max(dp[i][j], dp[i-1][j-weights[i-1]]+values[i-1])
				fmt.Printf("dp[%d][%d] = %d\n", i, j, dp[i][j])
			}
		}
	}
	// 出力
	fmt.Println(weights)
	fmt.Println(values)
	fmt.Println(dp)
	fmt.Println(dp[N][M])
}
