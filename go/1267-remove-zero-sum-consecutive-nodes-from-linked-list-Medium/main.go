/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
    dict := make(map[int]*ListNode)
    dummy := &ListNode{
        Val: 0,
        Next: head,
    }
    ptr := dummy
    sum := 0
    for ptr != nil {
        sum += ptr.Val
        if v, ok := dict[sum]; !ok {
            dict[sum] = ptr
        } else {
            del(v, sum, dict)
            v.Next = ptr.Next
        }
        ptr = ptr.Next
    }
    return dummy.Next
}

func del(node *ListNode, n int, dict map[int]*ListNode) {
    ptr := node.Next
    sum := n + ptr.Val
    for sum != n {
        delete(dict, sum)
        ptr = ptr.Next
        sum += ptr.Val
    } 
}