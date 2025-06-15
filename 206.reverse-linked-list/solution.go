/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var current, previous *ListNode = head, nil
    for current != nil {
        previous, current, current.Next = current, current.Next, previous
    }
    return previous
}