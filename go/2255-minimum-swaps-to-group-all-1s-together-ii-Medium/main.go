func minSwaps(nums []int) int {
    totalOnes := func() int {
        onesCount := 0
        for _, num := range nums {
            onesCount += num
        }
        return onesCount
    }()

    maxWindowOnes := func(windowSize int) int {
        windowOnes := 0
        maxWindowOnes := 0
        previousNum := 0
        j := 0
        for i := range nums {
            windowOnes -= previousNum
            for ; j - i < windowSize; j++ {
                windowOnes += nums[j % len(nums)]
            }
            maxWindowOnes = max(maxWindowOnes, windowOnes)
            previousNum = nums[i]
        }
        return maxWindowOnes
    }

    return totalOnes - maxWindowOnes(totalOnes)
}
