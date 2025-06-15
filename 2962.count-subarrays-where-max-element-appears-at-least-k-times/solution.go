func countSubarrays(nums []int, k int) int64 {
    maxNum := func() int {
        maxNum := nums[0]
        for _, num := range nums[1:] {
            maxNum = max(maxNum, num)
        }
        return maxNum
    }()

    return func() int64 {
        res := int64(0)
        count := 0
        left := 0
        for right := range nums {
            if nums[right] == maxNum {
                count++
            }
            for ; count >= k; left++ {
                if nums[left] == maxNum {
                    count--
                }
            }
            res += int64(left)
        }
        return res
    }()
}