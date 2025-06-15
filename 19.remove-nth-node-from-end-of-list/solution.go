/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    q := head
    for i := 0; i < n && q != nil; i++ {
        q = q.Next
    }
    if q == nil {
        return head.Next
    }
    p := head
    for q.Next != nil {
        p = p.Next
        q = q.Next
    } 
    p.Next = p.Next.Next
    return head
}