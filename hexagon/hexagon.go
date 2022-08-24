package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	var s string
	fmt.Fscanln(in, &n, &m)

	matrix := make([][]string, n)
	mb := make([][]bool, n)

	// Чтение
	for i := 0; i < n; i++ {
		fmt.Fscanln(in, &s)
		matrix[i] = strings.Split(s, "")
		mb[i] = make([]bool, m)
	}

	// Обработка
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == "." || mb[i][j] {
				continue
			}
			if matrix[i][j] == "/" || matrix[i][j] == "\\" {
				matrix[n-i][m-j-1], matrix[i][j] = matrix[i][j], matrix[n-i][m-j-1]
				mb[n-i][m-j-1] = true
			} else {
				matrix[n-i-1][m-j-1], matrix[i][j] = matrix[i][j], matrix[n-i-1][m-j-1]
				mb[n-i-1][m-j-1] = true
			}

		}
		// fmt.Println()
		// fmt.Println(i)
		// fmt.Println(mb)

	}

	// Вывод
	for i := 0; i < n; i++ {
		fmt.Println(strings.Join(matrix[i], ""))
	}
}
