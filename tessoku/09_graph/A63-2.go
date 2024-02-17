package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

// 割られる数がm, 割る数がnで，あまりを返す
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
	if key <= array[index] {
		return true
	} else {
		return false
	}
}

// 目的の値 target の index を返す (ない場合は -1)
func binarySearch(array []int, key int) int {
	ng := -1         //「index = 0」が条件を満たすこともあるので、初期値は -1
	ok := len(array) // 「index = a.size()-1」が条件を満たさないこともあるので、初期値は a.size()

	/* ok と ng のどちらが大きいかわからないことを考慮 */
	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2
		if isok(array, mid, key) {
			ok = mid
		} else {
			ng = mid
		}
	}
	// 現状の実装では条件を満たす値の最小値が返される
	return ok
}

// 配列の要素の始まりが1のときは下のように書く
// a = append([]int{0}, a...)
func GetNextNInt(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = fastio.GetNextInt()
	}
	return a
}

func GetNextNMInt(n int, m int) [][]int {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = GetNextNInt(m)
	}
	return a
}

func GetNextNInt64(n int) []int64 {
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = fastio.GetNextInt64()
	}
	return a
}

func GetNextNMInt64(n int, m int) [][]int64 {
	a := make([][]int64, n)
	for i := 0; i < n; i++ {
		a[i] = GetNextNInt64(m)
	}
	return a
}

func GetNextNFloat64(n int) []float64 {
	a := make([]float64, n)
	for i := 0; i < n; i++ {
		a[i] = fastio.GetNextFloat64()
	}
	return a
}

func GetNextNString(n int) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = fastio.Text()
	}
	return s
}

func GetNextLineChars() []byte {
	s := fastio.Text()
	b := []byte(s)

	return b
}

func GetNextNLineChars(n int) [][]byte {
	bs := make([][]byte, n)
	for i := 0; i < n; i++ {
		s := fastio.Text()
		bs[i] = []byte(s)
	}
	return bs
}

func ByteToString(b byte) string {
	s := string([]byte{b})
	return s
}

func bitAllSearch() {
	n := 3 // 独立な要素数

	// bitsがn個の要素の各パターンを表す(全てを選ばないパターンを考慮しないならbit初期値1でいい)
	for bits := 0; bits < (1 << uint64(n)); bits++ {
		// bitsの各要素についてのループ
		var index []int
		for i := 0; i < n; i++ {
			// bitsのi個目の要素の状態がonかどうかチェック
			// fmt.Println((bits >> uint64(i)) & 1)
			// メモ : (bits & (1 << i)) = 0 というのは、i の j ビット目（2^j の位）が 0 であるための条件。
			// 　　　頻出なので知っておくようにしましょう。
			if (bits>>uint64(i))&1 == 1 {
				// 何かしらの処理
				index = append(index, 1)
			} else {
				index = append(index, 0)
			}
		}
		fastio.Println(index)
	}

	/*
		   出力　n = 3
		    [0 0 0]
			[1 0 0]
			[0 1 0]
			[1 1 0]
			[0 0 1]
			[1 0 1]
			[0 1 1]
			[1 1 1]
	*/
}

// Replace character "from" to "to" in argument string.
// Return replaced string.
func replace(s string, from string, to string) string {
	// split string
	ss := strings.Split(s, "")
	for i, v := range ss {
		if v == from {
			ss[i] = to
		}
	}
	s = ""
	for _, v := range ss {
		s = s + v
	}
	return s
}

// dfsによって，グラフが連結かどうか探索する．
func dfs(crt int, g map[int][]int, visited []bool, prev map[int]int) {
	visited[crt] = true
	for i := 0; i < len(g[crt]); i++ {
		next := g[crt][i]
		if !visited[next] {
			prev[next] = crt
			dfs(next, g, visited, prev)
		}
	}
	return
}

func (i *FastIo) getGraph(n int) map[int][]int {
	g := map[int][]int{}

	var a, b int
	for i := 0; i < n; i++ {
		a = fastio.GetNextInt()
		b = fastio.GetNextInt()
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	return g
}

func swap(s string, a int, b int) string {
	ss := strings.Split(s, "")
	tmp := ss[a]
	ss[a] = ss[b]
	ss[b] = tmp

	ans := ""
	for _, v := range ss {
		ans += v
	}

	return ans
}

// 最大公約数 (GCD) を計算する関数
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 最小公倍数 (LCM) を計算する関数
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

type grid struct {
	grid [][]interface{}
}

func (g *grid) getGridValue(x, y int) (v interface{}) {
	return g.grid[y][x]
}

func (g *grid) setGridValue(x, y int, v interface{}) {
	g.grid[y][x] = v
}

type edge struct {
	to   int
	cost int
}

type vertex struct {
	visited  bool // 頂点が処理済みかどうか
	number   int  // 頂点番号
	distance int  // (ある点からの)最短距離
}

type graph struct {
	edges   [][]edge
	vertexs []vertex
}

func newGraph(n int) graph {
	g := graph{
		edges:   make([][]edge, n+1),
		vertexs: make([]vertex, n+1),
	}
	for i := 0; i <= n; i++ {
		g.vertexs[i] = vertex{
			visited:  false,
			number:   i,
			distance: math.MaxInt,
		}
	}
	return g
}

// 辺をグラフに追加する．重み無しグラフならcost = 1にする．
func (g *graph) appendEdge(from, to, cost int) {
	g.edges[from] = append(g.edges[from], edge{
		to:   to,
		cost: cost,
	})
}

// dfsでグラフ探索する．引数は探索開始頂点の番号.
// 探索終了時，探索で訪れることができた頂点はvisited=trueになる．
func (g *graph) dfs(crt int) {
	g.vertexs[crt].visited = true

	for i := 0; i < len(g.edges[crt]); i++ {
		next := g.edges[crt][i].to

		if !g.vertexs[next].visited {
			g.dfs(next)
		}
	}
	return
}

// bfsでグラフ探索する．引数は探索開始頂点の番号．
// 探索終了時，探索で訪れることのできた頂点はvisited=trueになり，distanceには開始頂点からの最短経路長が保存される．
func (g *graph) bfs(crt int) {
	var que []int
	que = append(que, 1)
	g.vertexs[1].distance = 0
	g.vertexs[1].visited = true
	var pos int
	for len(que) != 0 {
		pos = que[0]
		que = que[1:]
		for i := 0; i < len(g.edges[pos]); i++ {
			nextEdge := g.edges[pos][i]
			nextNum := nextEdge.to
			nextVertex := g.vertexs[nextNum]
			if nextVertex.visited {
				continue
			}
			g.vertexs[nextNum].visited = true
			g.vertexs[nextNum].distance = g.vertexs[pos].distance + nextEdge.cost

			que = append(que, nextVertex.number)
		}
	}
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
	m := fastio.GetNextInt()

	g := newGraph(n)
	var a, b int
	for i := 0; i < m; i++ {
		a = fastio.GetNextInt()
		b = fastio.GetNextInt()
		g.appendEdge(a, b, 1)
		g.appendEdge(b, a, 1)
	}

	// 幅優先探索
	g.bfs(1)

	for _, v := range g.vertexs {
		fastio.Println(v)
	}

}
