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
	n := fastio.GetNextInt()
	m := fastio.GetNextInt()
	t := fastio.GetNextInt()

	a := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		a[i] = fastio.GetNextInt()
	}

	bonusRoom := make([]int, n+1)
	for i := 0; i < m; i++ {
		roomnum := fastio.GetNextInt()
		bonus := fastio.GetNextInt()
		bonusRoom[roomnum] = bonus
	}
	// debug fastio.Println(n, m, t, a, bonusRoom)

	for roomnum := 1; roomnum < n; roomnum++ {
		t += bonusRoom[roomnum]
		t -= a[roomnum-1]
		if t <= 0 {
			fastio.Println("No")
			return
		}
	}
	fastio.Println("Yes")
}
