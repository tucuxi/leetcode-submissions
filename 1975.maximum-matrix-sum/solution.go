func maxMatrixSum(matrix [][]int) int64 {
    var absSum int64
    negativesCount := 0
    minAbsValue := math.MaxInt

    for _, r := range matrix {
        for _, c := range r {
            if c < 0 {
                negativesCount++
                c = -c
            }
            absSum += int64(c)
            minAbsValue = min(minAbsValue, c)
        }
    }

    if negativesCount % 2 == 0 {
        return absSum
    }
    return absSum - 2 * int64(minAbsValue)
}