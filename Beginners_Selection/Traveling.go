package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	slicet := make([]int, n)
	slicex := make([]int, n)
	slicey := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&slicet[i])
		if err != nil {
			panic(err)
		}
		_, err = fmt.Scan(&slicex[i])
		if err != nil {
			panic(err)
		}
		_, err = fmt.Scan(&slicey[i])
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < n; i++ {
		var dt int
		var dd int
		if i == 0 {
			dt = slicet[i]
			dd = slicex[i] + slicey[i]
		} else {
			dt = slicet[i] - slicet[i-1]
			diffx := math.Abs(float64(slicex[i] - slicex[i-1]))
			diffy := math.Abs(float64(slicey[i] - slicey[i-1]))
			dd = int(diffx + diffy)
		}

		if dt < dd {
			fmt.Println("No")
			return
		}

		if (dt-dd)%2 != 0 {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
