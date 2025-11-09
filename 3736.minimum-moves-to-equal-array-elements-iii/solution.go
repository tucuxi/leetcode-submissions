func minMoves(nums []int) int {
    sum := 0
    maxNum := 0
    for _, n :=  range nums {
        sum += n
        maxNum = max(n, maxNum)
    }
    return len(nums) * maxNum - sum
}