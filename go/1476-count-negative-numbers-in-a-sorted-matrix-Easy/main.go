func countNegatives(grid [][]int) int {
    count := 0
    last := len(grid[0]) - 1
    for _, row := range grid {
        for last >= 0 && row[last] < 0 {
            last--
        }
        count += len(row) - last - 1
    }
    return count
}

/*
   4  3  2 -1
   3  2  1 -1
   1  1 -1 -2
  -1 -1 -2 -3
*/