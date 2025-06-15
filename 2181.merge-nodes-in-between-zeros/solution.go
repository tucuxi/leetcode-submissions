/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    sentinel := &ListNode{}
    newTail := sentinel
    sum := 0
    for p := head; p != nil; p = p.Next {
        if p.Val == 0 && sum > 0 {
            newTail.Next = &ListNode{Val: sum}
            newTail = newTail.Next
            sum = 0
        } else {
            sum += p.Val
        }
    }
    return sentinel.Next
}