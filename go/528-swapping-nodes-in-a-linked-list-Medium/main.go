/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
    var kb *ListNode
    ke := head
    l := 0
    for p := head; p != nil; p = p.Next {
        l++
        if l == k {
            kb = p
            ke = head
        } else {
            ke = ke.Next
        }
    }
    kb.Val, ke.Val = ke.Val, kb.Val
    return head
}