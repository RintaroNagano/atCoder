package main

import (
	"fmt"
)

func main() {
	var n int
	var y int
	_, err := fmt.Scan(&n, &y)
	if err != nil {
		panic(err)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			k := n - i - j
			// fmt.Printf("%d %d %d %d\n", i, j, k, 10000*k+5000*j+1000*i
			if 10000*k+5000*j+1000*i == y {
				fmt.Printf("%d %d %d\n", k, j, i)
				return
			}
		}
	}

	fmt.Println("-1 -1 -1")
}
