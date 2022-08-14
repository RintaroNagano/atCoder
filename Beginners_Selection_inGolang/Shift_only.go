package main

import (
	"fmt"
)

func main() {
	var length int
	var count int
	_, err := fmt.Scan(&length)
	if err != nil {
		panic(err)
	}
	n := make([]int, length)
	for i := 0; i < length; i++ {
		_, err := fmt.Scan(&n[i])
		if err != nil {
			panic(err)
		}
	}

	for end := false; ; {
		for i := 0; i < length; i++ {
			if n[i]%2 == 0 {
				n[i] = n[i] / 2
			} else {
				end = true
				break
			}
		}
		if end == true {
			break
		}
		count++
	}
	fmt.Println(count)
}
