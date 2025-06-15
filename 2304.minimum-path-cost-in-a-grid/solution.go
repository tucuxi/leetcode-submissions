func minPathCost(grid [][]int, moveCost [][]int) int {
    m, n := len(grid), len(grid[0])
    minimumCostColumns := make([]int, n)
    for j := range n {
        minimumCostColumns[j] = grid[m-1][j]
    }
    tempCostColumns := make([]int, n)
    for i := m - 2; i >= 0; i-- {
        for j := range n {
            tempCostColumns[j] = math.MaxInt
            value := grid[i][j]
            for k, c := range moveCost[value] {
                tempCostColumns[j] = min(tempCostColumns[j], value + c + minimumCostColumns[k])
            }
        }
        minimumCostColumns, tempCostColumns = tempCostColumns, minimumCostColumns
    }
    res := math.MaxInt
    for _, firstRowCost := range minimumCostColumns {
        if firstRowCost < res {
            res = firstRowCost
        }
    }
    return res
}