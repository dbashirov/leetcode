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

	var cnt, a, b int
	fmt.Fscanln(in, &cnt)

	for i := 0; i < cnt; i++ {
		fmt.Fscanln(in, &a, &b)
		fmt.Fprintln(out, a+b)
	}

}
