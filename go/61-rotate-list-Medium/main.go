/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 {
        return head
    }
    l := length(head)
    if l == 0 {
        return head
    }
    i := k % l
    if i == 0 {
        return head
    }
    p := advance(head, l - i - 1)
    rh := p.Next
    p.Next = nil
    q := rh
    for q.Next != nil {
        q = q.Next
    }
    q.Next = head
    return rh
    
}

func length(ln *ListNode) int {
    l := 0
    for ln != nil {
        l++
        ln = ln.Next
    }
    return l
}

func advance(ln *ListNode, n int) *ListNode {
    for n > 0 {
        ln = ln.Next
        n--
    }
    return ln
}