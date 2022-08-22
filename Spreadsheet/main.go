package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var cntLn, n, cntK int

	in := bufio.NewReader(os.Stdin)
	cntLn = readInt(in)

	for i := 0; i < cntLn; i++ {

		in.ReadString('\n')

		sl := readArrInt(in)
		n = sl[0]
		matrix := make([][]int, n)
		for st := 0; st < n; st++ {
			matrix[st] = readArrInt(in)
		}
		cntK = readInt(in)
		c := readArrInt(in)
		// fmt.Printf("cntLn = %d\n", cntLn)
		// fmt.Printf("n = %d, m=%d\n", n, m)
		// fmt.Println(matrix)
		// fmt.Printf("k = %d\n", cntK)
		// fmt.Printf("c = %v\n", c)

		for k := 0; k < cntK; k++ {
			swap := true
			for swap {
				swap = false
				for st := 0; st < n-1; st++ {
					if matrix[st][c[k]-1] > matrix[st+1][c[k]-1] {
						swap = true
						matrix[st], matrix[st+1] = matrix[st+1], matrix[st]
					}
				}
			}
		}
		for _, sl := range matrix {
			for _, v := range sl {
				fmt.Printf("%d ", v)
			}
			fmt.Printf("\n")
		}
		fmt.Println()
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

func readArrInt(in *bufio.Reader) []int {
	numbs := readLineNumbs(in)
	arr := make([]int, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr[i] = val
	}
	return arr
}
