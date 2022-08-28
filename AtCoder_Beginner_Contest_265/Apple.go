package main

import (
	"fmt"
)

func main() {
	var x int
	var y int
	var n int
	var price int
	_, err := fmt.Scan(&x, &y, &n)
	if err != nil {
		panic(err)
	}

	// 三倍より安い
	if 3*x <= y {
		price = n * x
	}
	// 三倍よりtakai
	if 3*x > y {
		ny := n / 3
		nx := n % 3
		price = ny*y + nx*x
	}
	fmt.Println(price)
}
