package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str, err := readLine()
	if err != nil {
		panic(err)
	}
	words := []string{"dream", "dreamer", "erase", "eraser"}
	wlength := len(words)

	suffix := "         " // 7文字以上の任意の文字列, スライスの長さ以上の領域へのアクセスでエラーを避けるために導入
	str = Reverse(str) + suffix
	for j := 0; j < wlength; j++ {
		words[j] = Reverse(words[j])
	}

	var can bool
	can = true
	for i := 0; i < len(str)-len(suffix); {
		can2 := false
		for j := 0; j < wlength; j++ {
			if words[j] == str[i:i+len(words[j])] {
				i += len(words[j])
				can2 = true
				break
			}
		}
		if can2 == false {
			can = false
			break
		}
	}
	if can == false {
		fmt.Println("NO")
	}
	if can == true {
		fmt.Println("YES")
	}
}

func readLine() (string, error) {
	var rdr = bufio.NewReaderSize(os.Stdin, 1000000)
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			return "", e
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf), nil
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
