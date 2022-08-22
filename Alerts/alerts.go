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

	var cntN, cntQ, t, id, nAlerts int
	fmt.Fscanln(in, &cntN, &cntQ)

	alerts := make(map[int]int, cntQ)

	for q := 0; q < cntQ; q++ {
		fmt.Fscanln(in, &t, &id)
		if t == 1 {
			nAlerts++
			alerts[id] = nAlerts
			// if id == 0 {
			// 	for k := range alerts {
			// 		alerts[k] = nAlerts
			// 	}
			// }
		} else {
			if v, ok := alerts[id]; ok {
				if v < alerts[0] {
					fmt.Fprintln(out, alerts[0])
				} else {
					fmt.Fprintln(out, v)
				}
			} else if v, ok := alerts[0]; ok {
				fmt.Fprintln(out, v)
			} else {
				fmt.Fprintln(out, 0)
			}
		}
	}
}
