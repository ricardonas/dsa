

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {

	// First step: Get length and at the same time get the last node.
	l, tail := 1, head

	for tail.Next != nil {
		l += 1
		tail = tail.Next
	}

	k = k % l

	if k == 0 {
		return head
	}

	cur := head

	for i := 0; i < (l - k); i++ {
		// Node cujo ponteiro ao próximo deve ser quebrado.
		cur = cur.Next
	}

	// Obtém os próximos nó antes de quebrar o vínculo.
	newHead := cur.Next

	// Quebra o nó.
	cur.Next = nil

	// Redireciona o nó do último nó.
	tail.Next = head

	return newHead
}