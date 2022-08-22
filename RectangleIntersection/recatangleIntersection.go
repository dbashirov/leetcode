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

	var n int
	// var s string
	fmt.Fscanln(in, &n)
	cnt := make(map[int]int, n)

	m := make([][4]int, n)

	// Чтение
	for i := 0; i < n; i++ {
		for j := 0; j < 4; j++ {
			fmt.Fscan(in, &m[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {



			if m[i][0] < m[j][0] && m[i][1] < m[j][1] && m[i][2] > m[j][0] && m[i][3] > m[j][1]  {
				cnt[i]++
				cnt[j]++
			} else if m[i][0] < m[j][2] && m[i][1] < m[j][3] && m[i][2] > m[j][2] && m[i][3] > m[j][3] {
				cnt[i]++
				cnt[j]++
			} else if m[i][0] < m[j][2] && m[i][1] < m[j][1] && m[i][2] > m[j][2] && m[i][3] > m[j][1] {
				cnt[i]++
				cnt[j]++
			} else if m[i][0] < m[j][0] && m[i][1] < m[j][3] && m[i][2] > m[j][0] && m[i][3] > m[j][3] {
				cnt[i]++
				cnt[j]++

			} else if m[j][0] < m[i][0] && m[j][1] < m[i][1] && m[j][2] > m[i][0] && m[j][3] > m[i][1]  {
				cnt[i]++
				cnt[j]++
			} else if m[j][0] < m[i][2] && m[j][1] < m[i][3] && m[j][2] > m[i][2] && m[j][3] > m[i][3] {
				cnt[i]++
				cnt[j]++
			} else if m[j][0] < m[i][2] && m[j][1] < m[i][1] && m[j][2] > m[i][2] && m[j][3] > m[i][1] {
				cnt[i]++
				cnt[j]++
			} else if m[j][0] < m[i][0] && m[j][1] < m[i][3] && m[j][2] > m[i][0] && m[j][3] > m[i][3] {
				cnt[i]++
				cnt[j]++
			
			}
		}
		fmt.Fprintf(out, "%d ", cnt[i])
	}
}
