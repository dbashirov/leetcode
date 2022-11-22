package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type SingelList struct {
	head *ListNode
}

func (s *SingelList) add(v int) {
	ln := &ListNode{
		Val: v,
	}
	if s.head == nil {
		s.head = ln
	} else {
		current := s.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = ln
	}
}

func printListNode(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func swapPairs(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		cur.Val, cur.Next.Val = cur.Next.Val, cur.Val
		cur = cur.Next.Next
	}
	return head
}

func main() {
	arr := []int{1}
	s := &SingelList{}
	for _, v := range arr {
		s.add(v)
	}
	printListNode(s.head)
	r := swapPairs(s.head)
	printListNode(r)
	// swapPairs(s.head)
}
