// https://leetcode.com/problems/reverse-linked-list?envType=problem-list-v2&envId=xi4ci4ig

type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution 1
func reverseList(head *ListNode) *ListNode {
	var reversed *ListNode = nil

	currentNode := head

	for currentNode != nil {
		reversed = &ListNode{ Val: currentNode.Val, Next: reversed }
		currentNode = currentNode.Next
	}

	return reversed
}

// Solution 2 - Two pointers
func reverseList(head *ListNode) *ListNode {
	prev, curr := nil, head
	
	// Original:	1 -> 2 -> 3 -> 4 -> 5
	// Result:		5 -> 4 -> 3 -> 2 -> 1
	for curr != nil {
		// Store a copy to the ptr next before it got changed.
		temp := curr.Next

		// Change the current element to point to the previous.
		curr.Next = prev

		// Moves 'previous' pointer to the curret reading node.
		prev = curr

		// Move 'curr' pointer forward with the store variable.
		curr = temp 
	}

	return prev

}

