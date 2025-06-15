/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    var (
        sentinel = &ListNode{0, head}
        p = sentinel
        q = head
        r *ListNode = nil
    )
	for i := 1; i < left; i++ {
		p, q = q, q.Next
	}
	for i := left; i <= right; i++ {
		r, q, q.Next = q, q.Next, r
	}
	p.Next, p.Next.Next = r, q
	return sentinel.Next
}