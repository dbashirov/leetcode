package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	rez := 0
	for _, val := range nums {
		// fmt.Println(val)
		rez ^= val
	}
	return (rez)
}

func main() {
	nums := []int{7, 6, 4, 7, 3, 1, 6, 4, 1}
	fmt.Println(singleNumber(nums))
}
