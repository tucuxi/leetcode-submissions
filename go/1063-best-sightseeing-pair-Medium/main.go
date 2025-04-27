func maxScoreSightseeingPair(values []int) int {
    res := 0
    leftIndex := 0
    for rightIndex := 1; rightIndex < len(values); rightIndex++ {
        leftScore := values[leftIndex] + leftIndex - rightIndex
        res = max(res, leftScore + values[rightIndex])
        if values[rightIndex] > leftScore {
            leftIndex = rightIndex
        }
    }
    return res
}