package main

import (
	"fmt"
)

func main() {
	var n [2]int
	fmt.Scan(&n[0], &n[1])
	sum := n[0] * n[1]
	if sum%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
