package main

import (
	"fmt"
)

func main() {
	var length int
	var max int
	var secondMax int
	var alice int
	var bob int
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

	for len(slice) > 0 {
		max = 0
		secondMax = 0
		for i := 0; i < len(slice); i++ {
			if slice[i] > max {
				secondMax = max
				max = slice[i]
			} else {
				if slice[i] > secondMax {
					secondMax = slice[i]
				}
			}
		}
		for i := 0; i < len(slice); i++ {
			if slice[i] == max {
				slice = append(slice[:i], slice[i+1:]...)
				break
			}
		}
		for i := 0; i < len(slice); i++ {
			if slice[i] == secondMax {
				slice = append(slice[:i], slice[i+1:]...)
				break
			}
		}
		alice += max
		bob += secondMax
		// fmt.Println(max)
		// fmt.Println(secondMax)
	}
	//fmt.Println(alice)
	//fmt.Println(bob)
	fmt.Println(alice - bob)
}
