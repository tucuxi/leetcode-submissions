func gridGame(grid [][]int) int64 {
    var sum1, sum2, minSum int64
    minSum = math.MaxInt64

    for _, points := range grid[0] {
        sum1 += int64(points)
    }
    for i := range grid[0] {
        sum1 -= int64(grid[0][i])
        minSum = min(minSum, max(sum1, sum2))
        sum2 += int64(grid[1][i])
    } 

    return minSum
}