/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
    var values, stack []int

    for p := head; p != nil; p = p.Next {
        values = append(values, p.Val)
    }

    res := make([]int, len(values))

    for i, value := range values {
        for j := len(stack) - 1; j >= 0 && values[stack[j]] < value; j-- {
            res[stack[j]] = value
            stack = stack[:j]
        }
        stack = append(stack, i)
    }

    return res
}