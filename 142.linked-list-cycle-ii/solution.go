/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    p, q := head, head
    for q != nil && q.Next != nil {
        p, q = p.Next, q.Next.Next
        if p == q {
            p = head
            for p != q {
                p, q = p.Next, q.Next
            }
            return p
        }
    }
    return nil
}