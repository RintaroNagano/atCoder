package main

import (
	"fmt"
)

func main() {
	var yen [3]int
	var goalSum int
	var count int
	_, err := fmt.Scan(&yen[0], &yen[1], &yen[2], &goalSum)
	if err != nil {
		panic(err)
	}
	for i := 0; i <= yen[0]; i++ {
		for j := 0; j <= yen[1]; j++ {
			for k := 0; k <= yen[2]; k++ {
				total := i*500 + j*100 + 50*k
				if total == goalSum {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
