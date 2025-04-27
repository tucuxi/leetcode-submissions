/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
    carry := rec(head)
    if carry == 0 {
        return head
    }
    return &ListNode{carry, head}
}

func rec(node *ListNode) int {
    if node == nil {
        return 0
    }
    carry := rec(node.Next)
    v := 2*node.Val + carry
    node.Val = v % 10
    return v/10
}