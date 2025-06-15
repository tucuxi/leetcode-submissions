/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	sentinel := &ListNode{math.MinInt, head}
	prev, curr := sentinel, head
	for curr != nil {
		next := curr.Next
		if curr.Val < prev.Val {
			prev.Next = next
			insp := sentinel
			for insp.Next.Val <= curr.Val {
				insp = insp.Next
			}
			curr.Next = insp.Next
			insp.Next = curr
		} else {
			prev = curr
		}
		curr = next
	}
	return sentinel.Next
}