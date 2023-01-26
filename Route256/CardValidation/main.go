package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cnt, n, m int
	fmt.Fscanln(in, &cnt)

	for t := 0; t < cnt; t++ {
		fmt.Fscanln(in, &n, &m)
		matrix := make([][]byte, n)
		for i := 0; i < n; i++ {
			matrix[i] = make([]byte, m)
			for j := 0; j < m; j++ {
				matrix[i][j], _ = in.ReadByte()
			}
			in.ReadLine()
		}
		// fmt.Println(matrix)

		fmt.Fprintln(out, detour(&matrix, n, m))	
	}

}

func findAllElement(matrix *[][]byte, el byte, i, j, n, m int) {
	if i<0 || i>=n {
		return
	}
	if j < 0 || j >= m {
		return
	}

	if (*matrix)[i][j] == el {
		(*matrix)[i][j] = 0
		findAllElement(matrix, el, i, j+2, n, m)	
		findAllElement(matrix, el, i, j-2, n, m)

		findAllElement(matrix, el, i+1, j+1, n, m)
		findAllElement(matrix, el, i+1, j-1, n, m)

		findAllElement(matrix, el, i-1, j-1, n, m)
		findAllElement(matrix, el, i-1, j+1, n, m)
	}
}

func detour(matrix *[][]byte, n, m int) string {
	control := make(map[byte]struct{})
	for i:=0; i < n; i++ {
		for j :=0; j<m; j++ {
			if (*matrix)[i][j] == 0 || (*matrix)[i][j] == 46 {
				continue
			}
			if _, ok := control[(*matrix)[i][j]]; ok {
				return "NO"
			}
			control[(*matrix)[i][j]] = struct{}{}
			findAllElement(matrix, (*matrix)[i][j], i, j, n, m)
			// fmt.Println(*matrix)
		}
	}
	return "YES"
}
