package main

import "fmt"

func permute(nums []int) [][]int {
	count := len(nums)
	res := [][]int{nums}
	for i := 0; i < count; i++ {
		for _, sl := range res {
			temp := clone(sl)
			for j := i; j > 0; j-- {
				temp[j], temp[j-1] = temp[j-1], temp[j]
				res = append(res, temp)
				temp = clone(temp)
			}
		}
	}
	return res
}

func clone(sl []int) []int {
	// var c, python, java = true, false, string
	temp := make([]int, len(sl))
	// fmt.Println(c, python,java)
	copy(temp, sl)
	return temp
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
