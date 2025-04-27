func minOperations(grid [][]int, x int) int {
    elems := make([]int, 0, len(grid) * len(grid[0]))
    rest := grid[0][0] % x

    for _, row := range grid {
        for _, cell := range row {
            if cell % x != rest { return -1 }
            elems = append(elems, cell / x)
        }
    }

    slices.Sort(elems)

    res := 0
    i, j := 0, len(elems) - 1

    for i < j {
        if i < len(elems) - j - 1 {
            res += (i + 1) * (elems[i + 1] - elems[i])
            i++
        } else {
            res += (len(elems) - j) * (elems[j] - elems[j - 1])
            j--
        }
    }

    return res
}
