package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)

	// dp[i][j][k] は、区間 [i,j] において数字 k を出現回数 (区間内の数列を並べ替えたときの 2 倍に相当) 個使うときの、条件を満たす区間の個数
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, 10)
		}
	}

	// 初期値: 1 文字からなる区間は条件を満たす
	for i := 0; i < n; i++ {
		dp[i][i][int(s[i]-'0')] = 1
	}

	// 区間を広げていく
	for len := 2; len <= n; len++ {
		for i := 0; i+len-1 < n; i++ {
			j := i + len - 1
			for k := 0; k < 10; k++ {
				// 区間 [i,j-1] において数字 k を使う場合
				for l := i; l < j; l++ {
					for m := 0; m < 10; m++ {
						dp[i][j][k] += dp[i][l][m] * dp[l+1][j-1][k-m]
					}
				}
				// 区間 [i,j-1] において数字 k を使わない場合
				dp[i][j][k] += dp[i][j-1][k]
			}
		}
	}

	// 区間 [0,n-1] において、数字を出現回数 (区間内の数列を並べ替えたときの 2 倍に相当) 個使うときの条件を満たす区間の個数の合計を求める
	ans := 0
	for k := 0; k < 10; k++ {
		ans += dp[0][n-1][k]
	}

	fmt.Println(ans)
}
