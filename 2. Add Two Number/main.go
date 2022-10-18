package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func Ints2List(nums []int) *ListNode {
	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// var rezult []int

	carry := 0
	
	sum := new(ListNode)
	current := sum 

	for l1 != nil || l2 != nil || carry != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			v2, l2 = l2.Val, l2.Next
		}
		num := v1 + v2 + carry
		carry = num / 10
		current.Next = &ListNode{num % 10, nil}
		current = current.Next
	}
	
	return(sum.Next)
}

func main() {
	l1, l2 := []int{2,4,3}, []int{5,6,4}
	fmt.Println(l1, l2)
	fmt.Println(Ints2List(l1))
	fmt.Println(addTwoNumbers(Ints2List(l1), Ints2List(l2)))
}
