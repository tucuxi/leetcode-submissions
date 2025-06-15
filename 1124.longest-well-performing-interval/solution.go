func longestWPI(hours []int) int {
    result := 0
    runningSum := 0
    firstOccurrence := make(map[int]int)
    for i, h := range hours {
        if h > 8 {
            runningSum++
        } else {
            runningSum--
        }
        if runningSum > 0 {
            result = i + 1
            continue
        }
        if _, exists := firstOccurrence[runningSum]; !exists {
            firstOccurrence[runningSum] = i
        }
        if previousOccurrence, exists := firstOccurrence[runningSum - 1]; exists {
            result = max(result, i - previousOccurrence)
        }
    }
    return result
}