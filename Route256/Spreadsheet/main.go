package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	var cntLn, n, m, cntk, swapM int

	in := bufio.NewReader(os.Stdin)
	fmt.Fscanln(in, &cntLn)

	for i := 0; i < cntLn; i++ {

		in.ReadLine()

		fmt.Fscanln(in, &n, &m)
		matrix := make([][]int, n)
		for st := 0; st < n; st++ {
			matrix[st] = make([]int, m)
			for cl := 0; cl < m; cl++ {
				fmt.Fscan(in, &matrix[st][cl])
			}
			in.ReadLine()
		}
		// fmt.Println(matrix)

		fmt.Fscanln(in, &cntk)
		for k := 0; k < cntk; k++ {
			fmt.Fscan(in, &swapM)
			sort.SliceStable(matrix, func(ii, jj int) bool {
				return matrix[ii][swapM-1] < matrix[jj][swapM-1]
			})
		}
		in.ReadLine()
		for _, sl := range matrix {
			for _, v := range sl {
				fmt.Printf("%d ", v)
			}
			fmt.Println()
		}
		fmt.Println()
	}

}
