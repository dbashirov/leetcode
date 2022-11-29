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

func add(head *ListNode, v int) {
	ln := &ListNode{
		Val: v,
	}
	if head == nil {
		head = ln
	} else {
		current := head
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

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	for {
		added := false
		var min, count int
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if !added || min > lists[i].Val {
				min = lists[i].Val
				count = 1
				added = true
			} else if min == lists[i].Val {
				count++
			}
		}
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil || lists[i].Val != min{
				continue
			}
			lists[i] = lists[i].Next
		}
		if !added {
			break
		}
		for i:=count; i > 0; i-- {
			ln := &ListNode{
				Val: min,
			}
			if head == nil {
				head = ln
			} else {
				current := head
				for current.Next != nil {
					current = current.Next
				}
				current.Next = ln
			}
		}
	}
	return head
}

func main() {
	lists := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	arrListNode := []*ListNode{}

	for _, r := range lists {
		s := &SingelList{}
		for _, v := range r {
			s.add(v)
		}
		arrListNode = append(arrListNode, s.head)
	}
	for _, ln := range arrListNode {
		printListNode(ln)
	}

	printListNode(mergeKLists(arrListNode))

}
