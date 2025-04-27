func waysToSplitArray(nums []int) int {
    total := func() (res int64) {
        for _, num := range nums {
            res += int64(num)
        }
        return
    }()

    return func() (res int) {
        runningSum := int64(0)
        for _, num := range nums[:len(nums) - 1] {
            runningSum += int64(num)
            if runningSum >= total - runningSum {
                res++
            }
        }
        return
    }()
}