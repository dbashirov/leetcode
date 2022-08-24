package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	if len(prices) < 2 {
		return 0
	}

	min := prices[0]

	for _, value := range prices {
		if min > value {
			min = value
		}
		if value - min > profit {
			profit = value - min
		}
	}
	return(profit)
}

func main() {
	prices := []int{7,6,4,3,1}
	fmt.Println(maxProfit(prices))
}
