package main

import "fmt"

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if nums[start] > nums[mid] {
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if nums[mid] > target && target >= nums[start] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return -1
}

func main() {
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{1, 3}
	// nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
	target := 3
	fmt.Println(search(nums, target))
}
