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
	s := fastio.Text()
	s = " " + s
	one := "1"[0]
	row := make([]bool, 7)
	if s[1] == one {
		fastio.Println("No")
		return
	}
	if s[7] == one {
		row[0] = true
	}
	if s[4] == one {
		row[1] = true
	}
	if s[8] == one || s[2] == one {
		row[2] = true
	}
	if s[5] == one {
		row[3] = true
	}
	if s[9] == one || s[3] == one {
		row[4] = true
	}
	if s[6] == one {
		row[5] = true
	}
	if s[10] == one {
		row[6] = true
	}

	flag1 := false
	flag2 := false
	for i := 0; i < 7; i++ {
		if row[i] == true && !flag1 {
			flag1 = true
			continue
		}
		if row[i] == false && flag1 && !flag2 {
			flag2 = true
			continue
		}
		if row[i] && flag1 && flag2 {
			// fastio.Println(row)
			// fastio.Println(flag1, flag2)
			fastio.Println("Yes")
			return
		}
	}
	fastio.Println("No")
}
