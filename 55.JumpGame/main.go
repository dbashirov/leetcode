package main

import "fmt"

func canJump(nums []int) bool {
	// for i := 0; i < len(nums); {
	// 	if i+nums[i] >= len(nums)-1 {
	// 		return true
	// 	}
	// 	max, k := i+nums[i], i
	// 	for j := i + 1; j <= i+nums[i]; j++ {
	// 		if i+nums[j] >= len(nums) - 1 {
	// 			return true
	// 		}
	// 		if max < j+nums[j] {
	// 			max = j + nums[j]
	// 			k = j
	// 		}

	// 	}
	// 	if max == 0 || k == i {
	// 		return false
	// 	}
	// 	i = k
	// }
	// return false
	i := 0
	k := nums[i]
	for ; k > 0 && i <len(nums)-1; {
		i++
		k--
		if nums[i]>k{
			k = nums[i]
		}
	}
	return i >= len(nums)-1
}

func main() {
	// nums := []int{5,9,3,2,1,0,2,3,3,1,0,0}
	nums := []int{3,2,1,1,4}
	fmt.Println(canJump(nums))
}
