/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
    set := make(map[int]bool)
    for _, num := range nums {
        set[num]= true
    }
    sentinel := &ListNode{Next: head}
    for p := sentinel; p.Next != nil;  {
        if set[p.Next.Val] {
            p.Next = p.Next.Next
        } else {
            p = p.Next
        }
    }
    return sentinel.Next
}