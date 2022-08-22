package main

import "fmt"

func isValid(s string) bool {
	var sl = []string{}
	for i:=0; i<len(s); i++ {
		// fmt.Println(string(s[i]))
		if string(s[i]) == "(" || string(s[i]) == "{" || string(s[i]) == "[" {
			sl = append(sl, string(s[i]))
		} else if len(sl) == 0 {
			return false
		} else if string(s[i]) == ")" && string(sl[len(sl)-1]) != "(" {
			return false
		} else if string(s[i]) == "}" && string(sl[len(sl)-1]) != "{" {
			return false
		} else if string(s[i]) == "]" && string(sl[len(sl)-1]) != "[" {
			return false
		} else {
			sl = sl[:len(sl)-1]
		}
	}
	return len(sl) == 0
}

func main() {
	s := "{}"
	fmt.Println(isValid(s))
}