package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n [3]int
	var str string
	fmt.Scan(&n[1], &n[2], &n[0]) // データを格納する変数のアドレスを指定
	fmt.Scan(&str)
	sum := n[1] + n[2] + n[0]
	fmt.Println(strconv.Itoa(sum) + " " + str)
}
