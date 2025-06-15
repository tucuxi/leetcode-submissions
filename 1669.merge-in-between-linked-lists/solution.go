/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    p := list1
    for i := a; i > 1; i-- {
        p = p.Next
    }
    q := p.Next
    p.Next = list2
    for i := b - a; i > 0; i-- {
        q = q.Next
    }
    r := list2
    for r.Next != nil {
        r = r.Next
    }
    r.Next = q.Next
    return list1
}