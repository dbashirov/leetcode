package main

import "fmt"

// func combinationSum(candidates []int, target int) [][]int {
// 	result := [][]int{}
// 	for k, num := range candidates {
// 		if num == target {
// 			result = append(result, []int{num})
// 		} else if target%num == 0 {
// 			result = append(result, getSl(target/num, num))
// 		}
// 		sum := num
// 		for i := k + 1; i < len(candidates); i++ {
// 			if num+candidates[i] == target {
// 				result = append(result, []int{num, candidates[i]})
// 			} else {
// 				sum += candidates[i]
// 				if sum == target {
// 					result = append(result, []int{num, candidates[i]})
// 				}
// 			}
// 			if target-candidates[i] > 0 && target-candidates[i] != num && (target-candidates[i])%num == 0 {
// 				buf := getSl((target-candidates[i])/num, num)
// 				buf = append(buf, candidates[i])
// 				result = append(result, buf)
// 			}
// 		}

// 	}
// 	return result
// }

// func getSl(size, num int) []int {
// 	buf := make([]int, size)
// 	for i := 0; i < size; i++ {
// 		buf[i] = num
// 	}
// 	return buf
// }

func req(candidates [] int, result *[][] int, buf [] int, k int, target int) {
	
	if  target < 0 || k > len(candidates){
		return
	}

	if target == 0 {
		tbuf := make([]int, len(buf))
		copy(tbuf, buf)
		*result = append(*result, tbuf)	
	}

	for i:=k; i<len(candidates); i++ { 
		buf = append(buf, candidates[i])
		req(candidates, result, buf, i, target-candidates[i])
		buf = buf[:len(buf)-1]
	}
}

func reqCombinationSum(candidates [] int, target int) [][]int {
	result := [][]int{}
	
	req(candidates, &result, make([]int, 0), 0, target)
	
	return result
}

func main() {
	candidates := []int{2, 3, 6, 4, 7, 8}
	target := 7
	fmt.Println(reqCombinationSum(candidates, target))
}
