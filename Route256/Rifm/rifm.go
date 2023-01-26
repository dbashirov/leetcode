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

	var cntN, cntQ int
	var s string
	fmt.Fscanln(in, &cntN)
	dict := make([]string, cntN)
	for i := 0; i < cntN; i++ {
		fmt.Fscanf(in, "%s\n", &dict[i])
	}
	// fmt.Fprintln(out, dict)
	fmt.Fscanln(in, &cntQ)

	for q := 0; q < cntQ; q++ {
		fmt.Fscanf(in, "%s\n", &s)
		arr := make([]string, cntN)
		// buf := make([]string, cntN)
		for i := 0; i < len(dict); i++ {
			// fmt.Fprintln(out, 1, s)
			if s[len(s)-1] == dict[i][len(dict[i])-1] {
				arr = append(arr, dict[i])
				// fmt.Fprintln(out, s[len(s)], dict[i])
			}
		}
		fmt.Fprintln(out, arr)
		// for j := 1; j < len(s); j++ {
		// 	copy(buf, arr)
		// 	for i := 0; i < len(dict); i++ {
		// 		if s[len(s)-j] != dict[i][len(dict[i])-j] {
		// 			arr = append(arr)
		// 		}
		// 	}
		// }
	}

}
