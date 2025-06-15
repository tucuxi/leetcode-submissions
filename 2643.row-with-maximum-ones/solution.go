func rowAndMaximumOnes(mat [][]int) []int {
    maxCount, maxRow := -1, 0
    for i := range mat {
        count := 0
        for _, v := range mat[i] {
            if v == 1 {
                count++
            }
        }
        if count > maxCount {
            maxCount = count
            maxRow = i
        }
    }
    return []int{maxRow, maxCount}
}