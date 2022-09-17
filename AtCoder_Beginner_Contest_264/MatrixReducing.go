package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

// 配列の中で引数のint型の値と同じ数を持つ要素のインデックスが帰ってくる．
// 見つからなかった場合は配列長が帰ってくる
func binarySearch(array []int, target int) int {
	n := len(array)
	i := sort.Search(n, func(i int) bool {
		return array[i] >= target
	})
	return i
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
	var m1, m2 [11][11]int
	h1 := fastio.GetNextInt()
	w1 := fastio.GetNextInt()
	for i := 1; i <= h1; i++ {
		for j := 1; j <= w1; j++ {
			m1[i][j] = fastio.GetNextInt()
		}
	}
	h2 := fastio.GetNextInt()
	w2 := fastio.GetNextInt()
	for i := 1; i <= h2; i++ {
		for j := 1; j <= w2; j++ {
			m2[i][j] = fastio.GetNextInt()
		}
	}

	// fastio.Println(m1, m2)
	for bit := 1; bit < (1 << uint64(h1)); bit++ {
		var rowIndex []int
		for i := 0; i < h1; i++ {
			if (bit>>uint64(i))&1 == 1 {
				rowIndex = append(rowIndex, i+1)
			}
		}
		for bit2 := 1; bit2 < (1 << uint64(w1)); bit2++ {
			var colIndex []int
			for i := 0; i < w1; i++ {
				if (bit2>>uint64(i))&1 == 1 {
					colIndex = append(colIndex, i+1)
				}
			}
			// fastio.Println(rowIndex, colIndex)
			if len(rowIndex) != h2 || len(colIndex) != w2 {
				continue
			}

			//fastio.Println(rowIndex, colIndex)
			flag := false
			for i := 1; i <= h2; i++ {
				for j := 1; j <= w2; j++ {
					if m1[rowIndex[i-1]][colIndex[j-1]] != m2[i][j] {
						flag = true
					}
				}
			}

			if flag {
				continue
			} else {
				fastio.Println("Yes")
				return
			}
		}
	}

	fastio.Println("No")
}

// 10:56
