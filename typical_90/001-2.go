package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var fastio *FastIo

type FastIo struct {
	Scanner *bufio.Scanner
	Writer  *bufio.Writer
}

func NewFastIo(fp io.Reader, wfp io.Writer) *FastIo {
	const BufSize = 2000005
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, BufSize), BufSize)
	return &FastIo{Scanner: scanner, Writer: bufio.NewWriter(wfp)}
}
func (i *FastIo) Text() string {
	if !i.Scanner.Scan() {
		panic("scan failed")
	}
	return i.Scanner.Text()
}
func (i *FastIo) Atoi(s string) int                 { x, _ := strconv.Atoi(s); return x }
func (i *FastIo) GetNextInt() int                   { return i.Atoi(i.Text()) }
func (i *FastIo) Atoi64(s string) int64             { x, _ := strconv.ParseInt(s, 10, 64); return x }
func (i *FastIo) GetNextInt64() int64               { return i.Atoi64(i.Text()) }
func (i *FastIo) Atof64(s string) float64           { x, _ := strconv.ParseFloat(s, 64); return x }
func (i *FastIo) GetNextFloat64() float64           { return i.Atof64(i.Text()) }
func (i *FastIo) Print(x ...interface{})            { fmt.Fprint(i.Writer, x...) }
func (i *FastIo) Printf(s string, x ...interface{}) { fmt.Fprintf(i.Writer, s, x...) }
func (i *FastIo) Println(x ...interface{})          { fmt.Fprintln(i.Writer, x...) }

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func mod(m int, n int) int {
	if m < 0 {
		return m%n + n
	}
	return m % n
}

func abs(n int) int {
	if n > 0 {
		return n
	} else {
		return -n
	}
}

func abs64(n int64) int64 {
	if n > 0 {
		return n
	} else {
		return -n
	}
}

// indexが条件を満たすかどうか
func isok(array []int, index int, key int) bool {
	if array[index] <= key {
		return true
	} else {
		return false
	}
}

// 目的の値 target の index を返す (ない場合は -1)
func binarySearch(array []int, key int) int {
	ng := len(array) //「index = 0」が条件を満たすこともあるので、初期値は -1
	ok := -1         // 「index = a.size()-1」が条件を満たさないこともあるので、初期値は a.size()

	/* ok と ng のどちらが大きいかわからないことを考慮 */
	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2
		if isok(array, mid, key) {
			ok = mid
		} else {
			ng = mid
		}
	}
	/* ng は条件を満たさない最大の値、 okは条件を満たす最小の値になっている */
	return ok
}

func main() {
	fp := os.Stdin
	wfp := os.Stdout
	fastio = NewFastIo(fp, wfp)
	defer func() {
		fastio.Writer.Flush()
	}()
	solve()
}

func solve() {
	n := fastio.GetNextInt()
	l := fastio.GetNextInt64()
	k := fastio.GetNextInt()
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = fastio.GetNextInt64()
	}

	var crt int64
	var pre int64

	a = append(a, l)

	proc := func(score int64) bool {
		pre = 0
		crt = 0
		count := 0
		for i := 0; i < len(a); i++ {
			crt = a[i] - pre
			// fmt.Printf("crt: %d (a[i]=%d, pre=%d)\n", crt, a[i], pre)
			if crt >= score {
				count++
				// fmt.Println("crt: ", crt)
				// fmt.Println("count: ", count)
				pre = a[i]
			}
		}

		if k+1 <= count {
			// println(true)
			return true
		} else {
			// println(false)
			return false
		}
	}

	var ok int64
	ok = -1
	ng := l
	for abs64(ok-ng) > 1 {
		mid := (ok + ng) / 2
		// fmt.Println("mid", mid)
		if proc(mid) {
			ok = mid
			// fmt.Println(true)
		} else {
			ng = mid
			// fmt.Println(false)
		}
	}
	fastio.Println(ok)
}
