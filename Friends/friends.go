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
	fmt.Fscanln(in, &cntN, &cntM)

	friends := make(map[int][]int, cntN)

	for t := 0; t < cntM; t++ {
		var x, y int
		fmt.Fscanln(in, &x, &y)
		friends[x] = append(friends[x], y)
		friends[y] = append(friends[y], x)
	}

	// fmt.Fprintln(out, friends)

	for i := 1; i <= cntN; i++ {
		if _, ok := friends[i]; !ok {
			fmt.Fprintln(out, 0)
		} else {
			bufFriends := make(map[int]int)
			posFriends := make(map[int][]int)

			bf := make(map[int]struct{}, len(friends[i]))
			for _, v := range friends[i] {
				bf[v] = struct{}{}
			}
			// level := 0
			for _, v := range friends[i] {
				// countFriends(&friends, &posFriends, v, i, 1)
				for _, f := range friends[v] {
					if f != i {
						if _, ok := bf[f]; !ok {
							bufFriends[f]++
							posFriends[bufFriends[f]] = append(posFriends[bufFriends[f]], f)
						}
					}
				}
			}
			// fmt.Fprintln(out, posFriends)
			if len(posFriends) > 0 {
				keys := make([]int, 0, len(posFriends))
				for k := range posFriends {
					keys = append(keys, k)
				}
				sort.Ints(keys)
				sort.Ints(posFriends[keys[len(keys)-1]])

				for _, v := range posFriends[keys[len(keys)-1]] {
					fmt.Fprintf(out, "%d ", v)
				}
				fmt.Fprintf(out, "\n")
			} else {
				fmt.Fprintln(out, 0)
			}

		}

	}
}
