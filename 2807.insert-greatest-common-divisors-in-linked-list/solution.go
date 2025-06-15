/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    p := head
    for p != nil && p.Next != nil {
        q := p.Next
        p.Next = &ListNode{gcd(p.Val, q.Val), q}
        p = q
    }
    return head
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}