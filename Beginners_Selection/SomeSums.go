package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	var from int
	var to int
	var sum int
	var allSum int
	fmt.Scan(&n, &from, &to)

	for i := 1; i <= n; i++ {
		sum = 0
		str := strconv.Itoa(i)
		ss := strings.Split(str, "")
		for _, s := range ss {
			val, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			sum += val
		}
		if from <= sum && sum <= to {
			allSum += i
		}
	}
	fmt.Println(allSum)
}
