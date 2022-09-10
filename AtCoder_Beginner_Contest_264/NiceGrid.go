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
	i := fastio.GetNextInt()
	j := fastio.GetNextInt()

	var grid [][]string
	grid = make([][]string, 16)
	for i := 0; i < 16; i++ {
		grid[i] = make([]string, 16)
	}
	for i := 1; i <= 15; i++ {
		for j := 1; j <= 15; j++ {
			grid[i][j] = "white"
		}
	}
	// for i := 0; i < len(grid); i++ {
	// 	fastio.Println(grid[i])
	// }
	paintSquere := func(grid [][]string, edge int, row int) {
		col := row
		for i := 0; i < edge; i++ {
			grid[row][i+col] = "black"
			grid[row+edge-1][i+col] = "black"
			grid[i+row][col] = "black"
			grid[i+row][col+edge-1] = "black"
		}
	}
	paintSquere(grid, 15, 1)
	paintSquere(grid, 11, 3)
	paintSquere(grid, 7, 5)
	paintSquere(grid, 3, 7)

	// for i := 0; i < len(grid); i++ {
	// 	fastio.Println(grid[i])
	// }
	fastio.Println(grid[i][j])
}
