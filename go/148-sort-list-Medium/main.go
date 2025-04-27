/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    other := divide(head)
    head1 := sortList(head)
    head2 := sortList(other)
    return merge(head1, head2)
}

func divide(head *ListNode) *ListNode {
    slow := head
    fast := slow.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    newHead := slow.Next
    slow.Next = nil
    return newHead
}

func merge(p, q *ListNode) *ListNode {
    sentinel := &ListNode{}
    previous := sentinel
    for p != nil && q != nil {
        if p.Val < q.Val {
            previous.Next = p
            p = p.Next
        } else {
            previous.Next = q
            q = q.Next
        }
        previous = previous.Next
    }
    if p != nil {
        previous.Next = p
    } else {
        previous.Next = q
    }
    return sentinel.Next
}