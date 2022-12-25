package main

import (
	"fmt"
)

func main() {
	nodeList := []*ListNode{
		{
			Val: 2,
		},
		{
			Val: 4,
		},
		{
			Val: 3,
		},
	}
	nodeList[0].Next = nodeList[1]
	nodeList[1].Next = nodeList[2]
	l1 := nodeList[0]

	nodeList = []*ListNode{
		{
			Val: 5,
		},
		{
			Val: 6,
		},
		{
			Val: 4,
		},
	}
	nodeList[0].Next = nodeList[1]
	nodeList[1].Next = nodeList[2]
	l2 := nodeList[0]
	i := 0
	tmp := addTwoNumbers(l1, l2)
	for {
		if tmp != nil {
			fmt.Printf("i=%d, tmp=%+v\n", i, *tmp)
			tmp = tmp.Next
			i = i + 1
		} else {
			break
		}
	}

	nodeList = []*ListNode{
		{
			Val: 0,
		},
	}
	l1 = nodeList[0]

	nodeList = []*ListNode{
		{
			Val: 0,
		},
	}
	l2 = nodeList[0]
	i = 0
	tmp = addTwoNumbers(l1, l2)
	for {
		if tmp != nil {
			fmt.Printf("i=%d, tmp=%+v\n", i, *tmp)
			tmp = tmp.Next
			i = i + 1
		} else {
			break
		}
	}

	nodeList = []*ListNode{
		{
			Val: 9,
		},
		{
			Val: 9,
		},
		{
			Val: 9,
		},
		{
			Val: 9,
		},
		{
			Val: 9,
		},
		{
			Val: 9,
		},
		{
			Val: 9,
		},
	}
	nodeList[0].Next = nodeList[1]
	nodeList[1].Next = nodeList[2]
	nodeList[2].Next = nodeList[3]
	nodeList[3].Next = nodeList[4]
	nodeList[4].Next = nodeList[5]
	nodeList[5].Next = nodeList[6]
	l1 = nodeList[0]

	nodeList = []*ListNode{
		{
			Val: 9,
		},
		{
			Val: 9,
		},
		{
			Val: 9,
		},
		{
			Val: 9,
		},
	}
	nodeList[0].Next = nodeList[1]
	nodeList[1].Next = nodeList[2]
	nodeList[2].Next = nodeList[3]
	l2 = nodeList[0]
	i = 0
	tmp = addTwoNumbers(l1, l2)
	for {
		if tmp != nil {
			fmt.Printf("i=%d, tmp=%+v\n", i, *tmp)
			tmp = tmp.Next
			i = i + 1
		} else {
			break
		}
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tmp := head
	carry := 0
	for l1 != nil || l2 != nil {
		val := carry
		if l1 != nil {
			val = val + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val = val + l2.Val
			l2 = l2.Next
		}

		carry = val / 10
		val = val % 10
		tmp.Next = &ListNode{
			Val: val,
		}
		tmp = tmp.Next
	}
	if carry != 0 {
		tmp.Next = &ListNode{
			Val: carry,
		}
	}
	return head.Next
}
