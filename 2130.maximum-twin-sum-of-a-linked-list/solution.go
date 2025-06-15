/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    stack := make([]int, 0)
    for p:= head; p != nil; p = p.Next {
        stack = append(stack, p.Val)
    }
    max := 0
    k := len(stack) / 2
    p := head
    for i := len(stack) - 1; i >= k; i-- {
        if twin := p.Val + stack[i]; twin > max {
            max = twin
        }
        p = p.Next
    }
    return max
}