/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
    p := criticalPoints(head)
    if len(p) < 2 {
        return []int{-1, -1}
    }
    minDist := math.MaxInt
    for i := 1; i < len(p); i++ {
        minDist = min(minDist, p[i] - p[i - 1])
    }
    return []int{minDist, p[len(p) - 1] - p[0]}

}

func criticalPoints(head *ListNode) []int {
    var criticalPoints []int
    a := head.Val
    b := head.Next.Val
    p := head.Next.Next
    i := 1
    for p != nil {
        c := p.Val
        if isCritical(a, b, c) {
            criticalPoints = append(criticalPoints, i)
        }
        a, b = b, c
        p = p.Next
        i++
    }
    return criticalPoints
}

func isCritical(a, b, c int) bool {
    return b > a && b > c || b < a && b < c
}