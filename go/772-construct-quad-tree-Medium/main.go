/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
    return constructArea(grid, 0, 0, len(grid))
}

func constructArea(grid [][]int, row, col, n int) *Node {
    elem := grid[row][col]
    for i := row; i < row + n; i++ {
        for j := col; j < col + n; j++ {
            if elem != grid[i][j] {
                nHalf := n / 2
                return &Node{
                    IsLeaf: false,
                    TopLeft: constructArea(grid, row, col, nHalf),
                    TopRight: constructArea(grid, row, col + nHalf, nHalf),
                    BottomLeft: constructArea(grid, row + nHalf, col, nHalf),
                    BottomRight: constructArea(grid, row + nHalf, col + nHalf, nHalf),
                }
            }
        }
    }
    return &Node{Val: elem == 1, IsLeaf: true}
}