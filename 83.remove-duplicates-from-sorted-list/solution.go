/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    var p *ListNode
    for q := head; q != nil; q = q.Next {
        if p == nil || p.Val != q.Val {
            p = q
        } else {
            p.Next = q.Next
        }
    }
    return head
}