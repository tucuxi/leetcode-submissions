/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    n := 0
    for p := head; p != nil; p = p.Next {
        n++
    }

    st := make([]int, 0, n/2)
    p := head

    for range n/2 {
        st = append(st, p.Val)
        p = p.Next
    }
    
    res := 0

    for i := n/2 - 1; i >=0; i-- {
        t := st[i]
        res = max(res, p.Val + t)
        p = p.Next
    }

    return res
}