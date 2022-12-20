package main

import "fmt"

func rotate(nums []int, k int) {
	kl := k % len(nums)
	move(&nums, 0, len(nums)-1)
	move(&nums, 0, kl-1)
	move(&nums, kl, len(nums)-1)
}

func move(nums *[]int, begin, end int) {
	for begin < end {
		(*nums)[begin], (*nums)[end] = (*nums)[end], (*nums)[begin]
		begin++
		end--
	}
}

func main() {
	var (
		nums = []int{1, 2, 3, 4, 5, 6, 7}
		k    = 4
	)
	rotate(nums, k)
	fmt.Println(nums)
}
