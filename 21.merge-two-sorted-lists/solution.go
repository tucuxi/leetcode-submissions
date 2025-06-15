/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    sentinel := &ListNode{}
    pred := sentinel
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            pred.Next = list1
            pred = list1
            list1 = list1.Next
        } else {
            pred.Next = list2
            pred = list2
            list2 = list2.Next
        }
    }
    if list1 == nil {
        pred.Next = list2
    } else {
        pred.Next = list1
    }
    return sentinel.Next
}