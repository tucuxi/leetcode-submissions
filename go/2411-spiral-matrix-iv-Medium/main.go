/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {
    res := make([][]int, m)
    for i := range res {
        res[i] = make([]int, n)
        for j := range res[i] {
            res[i][j] = -1
        }
    }
    steps := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    dist := []int{n, m - 1}
    dir := 0
    row := 0
    col := -1
    cur := dist[0]
    for p := head; p != nil; p = p.Next {
        if cur == 0 {
            dist[dir % 2]--
            dir = (dir + 1) % len(steps)
            cur = dist[dir % 2]
        }
        row += steps[dir][0]
        col += steps[dir][1]
        res[row][col] = p.Val
        cur--
    }
    return res
}