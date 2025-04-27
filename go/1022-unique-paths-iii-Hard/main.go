const (
    obstacle       = -1
    emptySquare    =  0
    startingSquare =  1
    endingSquare   =  2
    visited        =  3
)

var (
    steps = [][]int{{ -1, 0 }, { 1, 0 }, { 0, -1 }, { 0, 1 }}
)

func uniquePathsIII(grid [][]int) int {
    paths := 0
    maxLength := len(grid) * len(grid[0])
    startRow, startCol := 0, 0
    for i := range grid {
        for j := range grid[i] {
            if square := grid[i][j]; square == obstacle {
                maxLength--
            } else if square == startingSquare {
                startRow, startCol = i, j
            }
        }
    }

    var dfs func(int, int, int)
    dfs = func(length, row, col int) {
        for _, step := range steps {
            r, c := row + step[0], col + step[1]
            if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
                continue
            }
            if square := grid[r][c]; square == endingSquare {
                if length + 1 == maxLength {
                    paths++
                }
            } else if square == emptySquare {
                grid[r][c] = visited
                dfs(length + 1, r, c)
                grid[r][c] = square
            }
        }
    }

    dfs(1, startRow, startCol)
    return paths
}