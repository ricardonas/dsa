/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    p1, p2 := headA, headB

    if p1 == nil || p2 == nil {
        return nil
    }

    // A troca de listas resolve o problema do tamanho diferente.
    // Precisa considerar a posição nil ao final das listas, também é uma posição válida.
    // Se as listas não tiverem intersecção, então, após a troca de lista, os ponteiros se encontrarão na posição nil.

    for p1 != p2 {

        if p1 != nil {
            p1 = p1.Next
        } else {
            p1 = headB
        }

        if p2 != nil {
            p2 = p2.Next
        } else {
            p2 = headA
        }
    }

    return p1
}
