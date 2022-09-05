package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var iost *Iost

type Iost struct {
	Scanner *bufio.Scanner
	Writer  *bufio.Writer
}

func NewIost(fp io.Reader, wfp io.Writer) *Iost {
	const BufSize = 2000005
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, BufSize), BufSize)
	return &Iost{Scanner: scanner, Writer: bufio.NewWriter(wfp)}
}
func (i *Iost) Text() string {
	if !i.Scanner.Scan() {
		panic("scan failed")
	}
	return i.Scanner.Text()
}
func (i *Iost) Atoi(s string) int                 { x, _ := strconv.Atoi(s); return x }
func (i *Iost) GetNextInt() int                   { return i.Atoi(i.Text()) }
func (i *Iost) Atoi64(s string) int64             { x, _ := strconv.ParseInt(s, 10, 64); return x }
func (i *Iost) GetNextInt64() int64               { return i.Atoi64(i.Text()) }
func (i *Iost) Atof64(s string) float64           { x, _ := strconv.ParseFloat(s, 64); return x }
func (i *Iost) GetNextFloat64() float64           { return i.Atof64(i.Text()) }
func (i *Iost) Print(x ...interface{})            { fmt.Fprint(i.Writer, x...) }
func (i *Iost) Printf(s string, x ...interface{}) { fmt.Fprintf(i.Writer, s, x...) }
func (i *Iost) Println(x ...interface{})          { fmt.Fprintln(i.Writer, x...) }
func isLocal() bool                               { return os.Getenv("WITHIN") == "TEMPTATION" }
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	if isLocal() {
		fp, _ = os.Open(os.Getenv("REMEMBER_YOU_AS_LONG_AS_I_CAN"))
	}
	iost = NewIost(fp, wfp)
	defer func() {
		iost.Writer.Flush()
	}()
	solve()
}
func solve() {
	n := iost.GetNextInt()
	m := iost.GetNextInt()
	aa := make([]int, n)
	for i := 0; i < n; i++ {
		aa[i] = iost.GetNextInt()
	}
	dp := makeGrid(n+3, m+3)
	const Inf = -100 //math.MinInt64
	for i := 0; i < n+3; i++ {
		for j := 0; j < m+3; j++ {
			dp[i][j] = Inf
		}
	}
	iost.Println(dp)
	dp[0][0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dp[i][j] == Inf {
				continue
			}
			dp[i+1][j] = max(dp[i+1][j], dp[i][j])
			dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j]+aa[i]*(j+1))
			iost.Println(dp)
		}
	}
	ans := dp[0][m]
	for i := 0; i <= n; i++ {
		ans = max(ans, dp[i][m])
		iost.Println(ans)
	}
	iost.Println(ans)
}
func makeGrid(h, w int) [][]int {
	index := make([][]int, h)
	data := make([]int, h*w)
	for i := 0; i < h; i++ {
		index[i] = data[i*w : (i+1)*w]
	}
	return index
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int { return -max(-a, -b) }
func abs(a int) int    { return max(a, -a) }
