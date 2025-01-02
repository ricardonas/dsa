package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var reversed *ListNode = nil

	currentNode := head

	for currentNode != nil {
		reversed = &ListNode{ Val: currentNode.Val, Next: reversed }
		currentNode = currentNode.Next
	}

	return reversed
}

