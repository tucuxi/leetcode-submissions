/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head.Next == nil {
        return nil
    }
    p, q := head, head
    for {
        if q.Next != nil {
            q = q.Next
        }
        if q.Next != nil {
            q = q.Next
        }
        if q.Next == nil {
            break
        }
        p = p.Next
    }
    p.Next = p.Next.Next
    return head
}