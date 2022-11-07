package main

import (
	"fmt"
)

func isPalindrome(x int) bool {

	// strs := strconv.Itoa(x)
	// if len(strs) <= 1 {
	// 	return(true)
	// }

	// for beg, end := 0, len(strs)-1; beg < end; beg, end = beg + 1, end - 1 {
	// 	if strs[beg] != strs[end] {
	// 		return(false)
	// 	}
	// }

	// return(true)

	if x < 0 || (x%10 == 0 && x != 0) {
		return (false)
	}

	revNumber := 0
	for x > revNumber {
		revNumber = revNumber*10 + (x % 10)
		x /= 10
	}

	return x == revNumber || x == revNumber/10
}

func main() {
	fmt.Println(isPalindrome(1221))
}
