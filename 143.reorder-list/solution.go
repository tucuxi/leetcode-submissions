/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    var prev *ListNode
    for slow != nil {
        prev, slow, slow.Next = slow, slow.Next, prev
    }

    first := head
    for prev.Next != nil {
        first.Next, first = prev, first.Next
        prev.Next, prev = first, prev.Next
    }
}