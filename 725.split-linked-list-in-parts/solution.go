/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
    n := 0
    for p := head; p != nil; p = p.Next {
        n++
    }
    res := make([]*ListNode, k)
    p := head
    for i := range res {
        res[i] = p
        l := n / k
        if i < n % k {
            l++
        }
        for ; l > 1; l-- {
            p = p.Next
        }
        if p == nil {
            break
        }
        p.Next, p = nil, p.Next
    }
    return res
}