package main

import (
	"bufio"
	"fmt"
	"math"
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
		// fmt.Println(arr)

		for i := 0; i < len(arr)-1; i++ {
			if arr[i] != -1 {
				fmt.Fprintf(out, "%d ", i+1)
				// arr[i] = -1
				min := float64(-1)
				minj := i + 1
				for j := i + 1; j < len(arr); j++ {
					if arr[j] == -1 {
						continue
					}
					// diff := arr[i] - arr[j]
					buf := math.Abs(float64(arr[i] - arr[j]))
					// fmt.Fprintf(out, "buff=%v min=%v\n", buf, min)
					if min == -1 || min > buf {
						minj = j
						min = buf
					}
				}
				arr[minj] = -1
				// fmt.Println(arr)
				fmt.Fprintf(out, "%d\n", minj+1)
			}
		}
		fmt.Fprintln(out)
	}
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
