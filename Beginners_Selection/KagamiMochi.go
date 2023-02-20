package main

import (
	"fmt"
	"sort"
)

func main() {
	var length int
	_, err := fmt.Scan(&length)
	if err != nil {
		panic(err)
	}
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		_, err := fmt.Scan(&slice[i])
		if err != nil {
			panic(err)
		}
	}
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	var pre int
	var count int
	pre = -1
	for i := 0; i < len(slice); i++ {
		if slice[i] != pre {
			count++
			pre = slice[i]
		}
	}
	fmt.Println(count)
}
