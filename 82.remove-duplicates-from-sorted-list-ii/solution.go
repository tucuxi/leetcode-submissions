/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    if head.Val == head.Next.Val && head.Next.Next == nil {
        return nil
    }
    var a, b, c *ListNode = nil, head, head.Next
    for c != nil {
        if b.Val == c.Val {
            for c != nil && b.Val == c.Val {
                c = c.Next
            }
            if a == nil {
                head = c
            } else {
                a.Next = c
            }
            b = c
            if c != nil {
                c = c.Next
            }
        } else {
            a, b, c = b, c, c.Next
        }
    }
    return head
}