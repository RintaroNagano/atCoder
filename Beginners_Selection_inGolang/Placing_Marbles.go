package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string
	var sum int
	sum = 0
	fmt.Scan(&str)
	ss := strings.Split(str, "")
	for _, s := range ss {
		val, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		sum += val
	}
	fmt.Println(sum)
}
