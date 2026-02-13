/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    sentinel := &ListNode{}
    current := sentinel
    carry := 0

    for l1 != nil || l2 != nil || carry > 0 {
        v1 := 0
        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        }

        v2 := 0
        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }
        
        sum := v1 + v2 + carry
        carry = sum / 10
        current.Next = &ListNode{Val: sum % 10}
        current = current.Next
    }

    return sentinel.Next
}