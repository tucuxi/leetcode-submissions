func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
    return max(slide(nums, firstLen, secondLen), slide(nums, secondLen, firstLen))
}

func slide(nums []int, firstLen, secondLen int) int {
    sum1, sum2 := 0, 0
    for i := range firstLen {
        sum1 += nums[i]
    }
    for i := range secondLen {
        sum2 += nums[firstLen + i]
    }

    maxSum1 := sum1
    res := sum1 + sum2
    for i := firstLen + secondLen; i < len(nums); i++ {
        sum1 += nums[i - secondLen] - nums[i - secondLen - firstLen]
        maxSum1 = max(maxSum1, sum1)
        sum2 += nums[i] - nums[i - secondLen]
        res = max(res, maxSum1 + sum2)
    }

    return res
}