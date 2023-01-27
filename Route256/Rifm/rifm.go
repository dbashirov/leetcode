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

	var cntN, cntQ, max, imax int
	var s string
	
	fmt.Fscanln(in, &cntN)
	dict := make([]string, cntN)
	for i := 0; i < cntN; i++ {
		fmt.Fscanf(in, "%s\n", &dict[i])
	}

	fmt.Fscanln(in, &cntQ)
	for q := 0; q < cntQ; q++ {
		fmt.Fscanf(in, "%s\n", &s)
		findMax := make(map[int]int, cntN)
		max = 0

		for i := 0; i < cntN; i++ {
			for j := 0; j < len(s); j++ {
			
				if s == dict[i]{
					break
				}
				if j >= len(dict[i]) {
					break
				}
				if s[len(s)-j-1] != dict[i][len(dict[i])-j-1] {
					break
				}

				findMax[i]++
			}
			if max < findMax[i] {
				imax = i
				max = findMax[i]
			}
		}

		if max > 0 {
			fmt.Fprintln(out, dict[imax])
		} else {
			fmt.Fprintln(out, dict[0])
		}
	}

}
