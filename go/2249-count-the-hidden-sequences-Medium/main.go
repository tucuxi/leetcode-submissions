func numberOfArrays(differences []int, lower int, upper int) int {
    num, minNum, maxNum := 0, 0, 0

    for _, d := range differences {
        num += d
        minNum = min(num, minNum)
        maxNum = max(num, maxNum)
    }

    return max(0, (upper - lower) - (maxNum - minNum) + 1)
}