package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Buffer(make([]byte, 500000*1+1), 500000)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
	}
	m := make(map[int]bool)
	for i := range a {
		m[a[i]] = true
	}
	p := make([][]int, 0, n)
	for i := range m {
		p = append(p, []int{i, 0})
	}
	sort.Slice(p, func(i, j int) bool { return p[i][0] > p[j][0] })
	for i := range p {
		p[i][1] = i
	}
	mm := make(map[int]int)
	for i := range p {
		mm[p[i][0]] = p[i][1]
	}
	ans := make([]int, n)
	for i := range a {
		ans[mm[a[i]]]++
	}
	for i := range ans {
		fmt.Println(ans[i])
	}
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
