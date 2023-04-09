package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(12345) // 乱数のシードを固定
	for i := 0; i < 10; i++ {

		N := 5                            // 2~200000のランダムな値
		X := rand.Intn(2000001) - 1000000 // -1000000~1000000のランダムな値
		A := make([]int, N)
		for i := 0; i < N; i++ {
			A[i] = X + rand.Intn(2*N+1) - N // Xを中心にランダムな値を生成
		}
		// テストケースの出力
		fmt.Println(N, X)
		for _, a := range A {
			fmt.Print(a, " ")
		}
		fmt.Println("")
	}
}
