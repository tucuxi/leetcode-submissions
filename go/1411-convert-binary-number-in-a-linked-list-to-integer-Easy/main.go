/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    res := 0
    for n := head; n != nil; n = n.Next {
        res = (res << 1) | n.Val
    }
    return res
}