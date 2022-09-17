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
	var grid [501][501]string
	h := fastio.GetNextInt()
	w := fastio.GetNextInt()

	for i := 1; i <= h; i++ {
		str := fastio.Text()
		for j := 1; j <= w; j++ {
			grid[i][j] = str[j-1 : j]
		}
	}

	i := 1
	j := 1
	m := map[string][2]int{
		"U": {-1, 0},
		"D": {1, 0},
		"L": {0, -1},
		"R": {0, 1},
	}
	var visited [501][501]bool
	for {
		visited[i][j] = true
		i += m[grid[i][j]][0]
		j += m[grid[i][j]][1]
		if i < 1 {
			i++
			break
		}
		if i > h {
			i--
			break
		}
		if j < 1 {
			j++
			break
		}
		if j > w {
			j--
			break
		}
		if visited[i][j] {
			fastio.Println("-1")
			return
		}
	}
	fastio.Println(i, j)
}

// 10:07
// 10:45
