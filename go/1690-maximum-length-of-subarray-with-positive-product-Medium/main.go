func getMaxLen(nums []int) int {
    res := 0
    startIndex := 0
    negativeIndex := -1
    even := true
    for i, num := range nums {
        if num == 0 {
            startIndex = i + 1
            negativeIndex = -1
            even = true
            continue
        }
        if num < 0 {
            even = !even
            if negativeIndex == -1 {
                negativeIndex = i
            }
        }
        if even {
            res = max(res, i - startIndex + 1)
        } else {
            res = max(res, i - negativeIndex) 
        }
    }
    return res
}