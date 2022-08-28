package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

var h int
var w int

func solve() {
	// uketori
	h = fastio.GetNextInt()
	w = fastio.GetNextInt()
	cs := make([][]string, h+1)

	for i := 1; i <= h; i++ {
		str := "a" + fastio.Text()
		cs[i] = strings.Split(str, "")
	}
	visitedPos := make([][]bool, 1000, 1000)
	for i := 0; i < len(visitedPos); i++ {
		visitedPos[i] = make([]bool, 1000)
	}
	i := 1
	j := 1
	for {
		visitedPos[i][j] = true
		c := cs[i][j]
		canMove := move(c, &i, &j)
		if !canMove {
			break
		}
		if visitedPos[i][j] {
			fastio.Println("-1")
			return
		}
	}
	fastio.Println(i, j)
}

func move(c string, i *int, j *int) bool {
	switch c {
	case "U":
		if *i != 1 {
			*i -= 1
			return true
		} else {
			return false
		}
	case "D":
		if *i != h {
			*i += 1
			return true
		} else {
			return false
		}

	case "L":
		if *j != 1 {
			*j -= 1
			return true
		} else {
			return false
		}

	case "R":
		if *j != w {
			*j += 1
			return true
		} else {
			return false
		}

	default:
		panic("panic: unexpected charactor")

	}
}
