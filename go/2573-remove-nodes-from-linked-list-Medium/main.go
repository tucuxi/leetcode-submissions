/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    head.Next = removeNodes(head.Next)
    if head.Val < head.Next.Val {
        return head.Next
    }
    return head
}
