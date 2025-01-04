// https://leetcode.com/problems/reverse-linked-list?envType=problem-list-v2&envId=xi4ci4ig

type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution 1
// Time: O(n)
// Space: O(n)
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
// Time: O(n)
// Space: O(n)
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		nextCopy := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextCopy
	}

	return prev
} 

