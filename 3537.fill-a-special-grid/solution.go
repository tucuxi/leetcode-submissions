func specialGrid(n int) [][]int {
    m := 1 << n
    res := make([][]int, m)
    for i := range res {
        res[i] = make([]int, m)
    }
    fill(res, 0, 0, m, 0)
    return res
}

func fill(arr [][]int, row, col, size, value int) {
    if size == 1 {
        arr[row][col] = value
    } else {
        half := size / 2
        q := half * half
        fill(arr, row, col + half, half, value) // top right
        fill(arr, row + half, col + half, half, value + q) // bottom right
        fill(arr, row + half, col, half, value + 2 * q) // bottom left
        fill(arr, row, col, half, value + 3 * q) // top left
    }
}