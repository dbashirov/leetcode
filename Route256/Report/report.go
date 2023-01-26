package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cntT int
	fmt.Fscanln(in, &cntT)

	for t := 0; t < cntT; t++ {
		var n int
		fmt.Fscanln(in, &n)

		arr := readArrInt(in)
		fmt.Fprintln(out, control(&arr))

	}
}

func control(arr *[]int) string {
	var last int
	report := make(map[int]struct{})

	for _, v := range *arr {
		if _, ok := report[v]; !ok {
			report[v] = struct{}{}
			last = v
		} else {
			if last != v {
				return "NO"
			}
		}
	}
	return "YES"
}

func readLineNumbs(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	return numbs
}

func readArrInt(in *bufio.Reader) []int {
	numbs := readLineNumbs(in)
	arr := make([]int, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr[i] = val
	}
	return arr
}
