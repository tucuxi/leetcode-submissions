type entry struct {
    cost, row, col int
}

func minimumObstacles(grid [][]int) int {
    directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

    cost := make([][]int, len(grid))
    for i := range cost {
        cost[i] = make([]int, len(grid[i]))
        for j := range cost[i] {
            cost[i][j] = math.MaxInt
        }
    }
    cost[0][0] = 0

    queue := list.New()
    queue.PushFront(&entry{0, 0, 0})

    for queue.Len() > 0 {
        current := queue.Front()
        queue.Remove(current)
        currentEntry := current.Value.(*entry)
        for _, d := range directions {
            r := currentEntry.row + d[0]
            c := currentEntry.col + d[1]
            if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || cost[r][c] < math.MaxInt {
                continue
            }
            if grid[r][c] == 1 {
                cost[r][c] = currentEntry.cost + 1
                queue.PushBack(&entry{cost[r][c], r, c})
            } else {
                cost[r][c] = currentEntry.cost
                queue.PushFront(&entry{cost[r][c], r, c})
            }
        }
    }
    return cost[len(cost) - 1][len(cost[0]) - 1]
}