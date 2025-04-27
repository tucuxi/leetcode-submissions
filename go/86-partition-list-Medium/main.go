/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    listL, listR := new(ListNode), new(ListNode)
    pL, pR := listL, listR
    for pCur := head; pCur != nil; {
        pNext := pCur.Next
        pCur.Next = nil
        if pCur.Val < x {
            pL.Next = pCur
            pL = pCur
        } else {
            pR.Next = pCur
            pR = pCur
        }
        pCur = pNext
    }
    pL.Next = listR.Next
    return listL.Next
}