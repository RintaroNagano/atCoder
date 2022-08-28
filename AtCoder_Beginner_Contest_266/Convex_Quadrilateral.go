package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
	ax := fastio.GetNextFloat64()
	ay := fastio.GetNextFloat64()
	bx := fastio.GetNextFloat64()
	by := fastio.GetNextFloat64()
	cx := fastio.GetNextFloat64()
	cy := fastio.GetNextFloat64()
	dx := fastio.GetNextFloat64()
	dy := fastio.GetNextFloat64()

	m := map[string][2]float64{}
	vector := func(tag string, ax float64, ay float64, bx float64, by float64) {
		m[tag] = [2]float64{bx - ax, by - ay}
	}
	vector("AB", ax, ay, bx, by)
	vector("BC", bx, by, cx, cy)
	vector("CD", cx, cy, dx, dy)
	vector("DA", dx, dy, ax, ay)
	crossProduct := func(ax float64, ay float64, bx float64, by float64) float64 {
		return ax*by - ay*bx
	}

	ABcp := crossProduct(m["AB"][0], m["AB"][1], m["BC"][0], m["BC"][1])
	BCcp := crossProduct(m["BC"][0], m["BC"][1], m["CD"][0], m["CD"][1])
	CDcp := crossProduct(m["CD"][0], m["CD"][1], m["DA"][0], m["DA"][1])
	DAcp := crossProduct(m["DA"][0], m["DA"][1], m["AB"][0], m["AB"][1])

	min1 := math.Min(ABcp, BCcp)
	min2 := math.Min(CDcp, DAcp)
	if math.Min(min1, min2) < 0 {
		fastio.Println("No")
		return
	}

	fastio.Println("Yes")
}
