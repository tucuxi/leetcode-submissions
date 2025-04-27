/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    if fast != nil {
        slow = slow.Next
    }

    var prev, curr *ListNode = nil, slow
    for curr != nil {
        prev, curr, curr.Next = curr, curr.Next, prev
    }

    left, right := head, prev
    for right != nil {
        if left.Val != right.Val {
            return false
        }
        left = left.Next
        right = right.Next
    }
    return true
}