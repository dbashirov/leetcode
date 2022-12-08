package main

import "fmt"

func longestValidParentheses(s string) int {
	
	l := 0
	st := []int{}
	st = append(st, -1)

	for i,v := range s {
		if v == '(' {
			st = append(st, i)
		} else {
			st = st[:len(st)-1]
			if len(st) == 0 {
				st = append(st, i)
			} else {
				l = max(l, i-st[len(st)-1])
			}
		}
	}
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "()(()"
	fmt.Println(longestValidParentheses(s))
}
