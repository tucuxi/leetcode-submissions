/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    nodes := make([]*ListNode, k)
    var previousGroupHead *ListNode
    groupHead := head
    groupIndex := 0
    for currentNode := head; currentNode != nil; {
        nodes[groupIndex] = currentNode
        currentNode = currentNode.Next
        groupIndex = (groupIndex + 1) % k
        if groupIndex > 0 {
            continue
        }
        // Finish group
        if previousGroupHead == nil {
            head = nodes[k - 1]
        } else {
            previousGroupHead.Next = nodes[k - 1]
        }
        previousGroupHead = groupHead
        groupHead = currentNode
        nodes[0].Next = currentNode
        for j := k - 1; j > 0; j-- {
            nodes[j].Next = nodes[j - 1]
        }
    }
    return head
}