/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var res *ListNode
    s1 := toStack(l1)
    s2 := toStack(l2)
    c := 0
    for len(s1) > 0 || len(s2) > 0 {
        v := c
        if len(s1) > 0 {
            v += s1[len(s1) - 1]
            s1 = s1[:len(s1) - 1]
        }
        if len(s2) > 0 {
            v += s2[len(s2) - 1]
            s2 = s2[:len(s2) - 1]
        }
        c = v / 10
        node := &ListNode{v % 10, res}
        res = node
    }
    if c > 0 {
        node := &ListNode{c, res}
        res = node
    }
    return res
}

func toStack(l *ListNode) []int {
    res := []int{}
    for l != nil {
        res = append(res, l.Val)
        l = l.Next
    }
    return res
}