/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    return sliceToBST(listToSlice(head))
}

func listToSlice(head *ListNode) []int {
    res := []int{}
    for p := head; p != nil; p = p.Next {
        res = append(res, p.Val)
    }
    return res
}

func sliceToBST(a []int) *TreeNode {
    if len(a) == 0 {
        return nil
    }
    m := len(a) / 2
    return &TreeNode{a[m], sliceToBST(a[:m]), sliceToBST(a[m + 1:])}
}