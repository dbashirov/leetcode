package main

import "fmt"

func subsets(nums []int) [][]int {
	rez := make([][]int, 1)
	for _, v := range nums {
		curSize := len(rez)
		for i := 0; i < curSize; i++ {
			buf := make([]int, len(rez[i]))
			copy(buf, rez[i])
			buf = append(buf, v)
			rez = append(rez, buf)
		}
	}
	// buf := make([]int, len(nums))
	// for i := 0; i < len(nums); i++ {
	// 	rez = append(rez, []int{nums[i]})
	// 	for j := i + 1; j < len(nums); j++ {
	// 		rez = append(rez, []int{nums[i], nums[j]})
	// 		if j-i > 1 {
	// 			rez = append(rez, nums[i:j+1])
	// 			if j < len(nums)-1 {
	// 				for jj := j + 1; jj < len(nums); jj++ {
	// 					rez = append(rez, append([]int{nums[i]}, nums[j:jj+1]...))
	// 				}
	// 			}
	// 			if i > 0 {
	// 				buf := make([]int, len(nums))
	// 				copy(buf, nums)
	// 				rez = append(rez, append(buf[:i+1], nums[j]))
	// 			}
	// 			if i > 0 && j < len(nums)-1 {
	// 				buf := make([]int, len(nums))
	// 				copy(buf, nums)
	// 				rez = append(rez, append(buf[:i+1], nums[j:]...))
	// 			}
	// 		}
	// 	}
	// }
	return rez
}

func main() {

	nums := []int{1, 2, 3, 4, 5}
	// nums := []int{9, 0, 3, 5, 7}
	fmt.Println(subsets(nums))

}

// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
