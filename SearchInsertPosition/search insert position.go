package main

import "fmt"

func searchInsert(nums []int, target int) int {

	low := 0
	high := len(nums) - 1
	possition := -1

	for low <= high {
		var mid int = (low + high) / 2
		guess := nums[mid]
		if guess == target {
			possition = mid
			break
		} else if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
		
	}
	if possition == -1 {
		possition = low		
	}
	return(possition)
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 5))
}
