package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cntN, cntM int
	var sum, l, t int64
	fmt.Fscanln(in, &cntN, &cntM)

	// arrServ := readArrInt(in)
	arrServ := make([]int64, cntN)
	for i := 0; i < cntN; i++ {
		fmt.Fscanf(in, "%d", &arrServ[i])
	}
	fmt.Fscanf(in, "\n")
	// sort.Ints(arrServ)
	sort.Slice(arrServ, func(i, j int) bool {
		return arrServ[i] < arrServ[j]
	})
	serv := make(map[int]int64, cntN)

	for i := 0; i < cntM; i++ {
		fmt.Fscanln(in, &t, &l)
		for j := 0; j < cntN; j++ {
			// fmt.Fprintln(out, serv[j], t, l)
			if v, ok := serv[j]; (ok && v <= t) || !ok {
				serv[j] = t + l
				sum += arrServ[j] * l
				// fmt.Fprintln(out, 1, arrServ[j], l, arrServ[j]*l)
				break
			}
		}
	}
	fmt.Fprintf(out, "%d\n", sum)
}
