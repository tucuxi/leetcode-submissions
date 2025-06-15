/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type entry struct {
    row int
    val int
}

func verticalTraversal(root *TreeNode) [][]int {
    cols := map[int][]entry{}
    minCol := math.MaxInt
    maxCol := math.MinInt
    
    var rec func(node *TreeNode, row, col int)
    rec = func(node *TreeNode, row, col int) {
        if node != nil {
            m := cols[col]
            if m == nil {
                m = []entry{}
            }
            cols[col] = append(m, entry{row, node.Val})
            if col < minCol {
                minCol = col
            }
            if col > maxCol {
                maxCol = col
            }
            rec(node.Left, row+1, col-1)
            rec(node.Right, row+1, col+1)
        }
    }
    
    rec(root, 0, 0)
    
    res := [][]int{}
    for i := minCol; i <= maxCol; i++ {
        entries := cols[i]
        sort.Slice(entries, func(i, j int) bool {
            return entries[i].row < entries[j].row || entries[i].row == entries[j].row && entries[i].val < entries[j].val
        })
        nodeVals := []int{}
        for _, v := range entries {
            nodeVals = append(nodeVals, v.val)
        }
        res = append(res, nodeVals)
    }
    return res
}