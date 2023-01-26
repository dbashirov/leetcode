package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "os"
	// "strconv"
)

func main() {
	// var cntLn, cntPr, buf, sum int
	var cntLn, sum int
	// fmt.Scanln(&cntLn)
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	cntLn = readInt(in)

	for i := 0; i < cntLn; i++ {
		readInt(in)
		prices := readMapInt(in)
		sum = 0
		for k, val := range prices {
			sum += (int(val/3)*2 + val%3) * k
		}
		fmt.Fprintln(out, sum)
	}
}

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}

func readLineNumbs(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	return numbs
}

func readMapInt(in *bufio.Reader) map[int]int {
	numbs := readLineNumbs(in)
	prices := make(map[int]int)
	for _, n := range numbs {
		val, _ := strconv.Atoi(n)
		prices[val]++
	}
	return prices
}
