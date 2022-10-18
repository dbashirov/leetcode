package main

import (
	"fmt"
)

func isBadVersion(version int) bool {
	return(version >= 9)
}

func firstBadVersion(n int) int {

	low := 1
	high := n

	for low <= high {
		var mid int = (low + high) / 2
		 
		if isBadVersion(mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return(low)
}

func main() {
	fmt.Println(firstBadVersion(20))
}
