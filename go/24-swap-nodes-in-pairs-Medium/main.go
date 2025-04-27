/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head != nil && head.Next != nil {
        head, head.Next, head.Next.Next = head.Next, swapPairs(head.Next.Next), head
    }
    return head
}