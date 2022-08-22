package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cntT int
	fmt.Fscanln(in, &cntT)

	for t := 0; t < cntT; t++ {
		var n, m int
		fmt.Fscanln(in, &n, &m)

		// Чтение в матрицу
		matrix := make([][]int, n)
		ships := []int{}

		for i := 0; i < n; i++ {
			line := readLine(in)
			matrix[i] = make([]int, m)
			for j := 0; j < m; j++ {
				if line[j] == 42 {
					matrix[i][j] = 1
				} else {
					matrix[i][j] = 0
				}
			}
		}
		// fmt.Fprintln(out, matrix)

		// Обработка матрицы
		rez := matrixProcessing(&matrix, &ships, n, m)
		fmt.Fprintln(out, rez)
		if rez == "YES" {
			sort.Ints(ships)
			for _, v := range ships {
				fmt.Fprintf(out, "%d ", v)
			}
			if len(ships) > 0 {
				fmt.Fprintf(out, "\n")
			}
		}
		// fmt.Fprintln(out, "После обработки:")
		// fmt.Fprintln(out, matrix)
	}
}

func matrixProcessing(matrix *[][]int, ships *[]int, n, m int) string {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (*matrix)[i][j] == 1 {
				iArr := make(map[int]int)
				jArr := make(map[int]int)
				if !itemProcessing(matrix, i, j, &iArr, &jArr) {
					return "NO"
				}

				// Контроль
				// fmt.Println(t, iArr)
				key := false
				iMax := 1
				for _, v := range iArr {
					if v > 1 {
						iMax = v
						if key {
							return "NO"
						}
						key = true
					}
				}
				key = false
				jMax := 1
				for _, v := range jArr {
					if v > 1 {
						jMax = v
						if key {
							return "NO"
						}
						key = true
					}
				}
				if iMax != jMax {
					return "NO"
				}
				*ships = append(*ships, iMax+jMax-1)
			}
		}
	}
	return "YES"
}

func itemProcessing(matrix *[][]int, i, j int, iArr, jArr *map[int]int) bool {
	if limits(matrix, i, j) {
		return true
	}

	if (*matrix)[i][j] == 1 {

		(*matrix)[i][j] = -1

		(*iArr)[i]++
		(*jArr)[j]++
		// fmt.Println(iArr)

		// Осматриваемся
		if !itemProcessing(matrix, i, j+1, iArr, jArr) { // Вправо
			return false
		}
		if !itemProcessing(matrix, i+1, j, iArr, jArr) { // Вниз
			return false
		}
		if !itemProcessing(matrix, i, j-1, iArr, jArr) { // Влево
			return false
		}
		if !itemProcessing(matrix, i-1, j, iArr, jArr) { // Вверх
			return false
		}
		// itemProcessing(matrix, i+1, j, iArr, jArr) // Вниз
		// itemProcessing(matrix, i, j-1, iArr, jArr) // Влево
		// itemProcessing(matrix, i-1, j, iArr, jArr) // Вверх

		// Проверка на букву Т
		if filed(matrix, i, j-1)+filed(matrix, i+1, j)+filed(matrix, i, j+1)+filed(matrix, i-1, j) > 2 {
			return false
		}

		// Диагонали
		if (filed(matrix, i+1, j+1) == 1 && filed(matrix, i, j+1)+filed(matrix, i+1, j) == 0) ||
			(filed(matrix, i-1, j+1) == 1 && filed(matrix, i, j+1)+filed(matrix, i-1, j) == 0) ||
			(filed(matrix, i-1, j-1) == 1 && filed(matrix, i-1, j)+filed(matrix, i, j-1) == 0) ||
			(filed(matrix, i+1, j-1) == 1 && filed(matrix, i, j-1)+filed(matrix, i+1, j) == 0) {
			// fmt.Println("2222: ", filed(matrix, i, j-1), i, j)
			return false
		}

		return true
	}

	return true
}

func filed(matrix *[][]int, i, j int) int {
	if limits(matrix, i, j) {
		return 0
	}

	if (*matrix)[i][j] == 1 || (*matrix)[i][j] == -1 {
		return 1
	}
	return 0
}

func limits(matrix *[][]int, i, j int) bool {
	if i < 0 || i >= len(*matrix) || j < 0 || j >= len((*matrix)[i]) {
		return true
	}
	return false
}

func readLine(in *bufio.Reader) string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	return line
}
