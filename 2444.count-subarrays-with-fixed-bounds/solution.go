func countSubarrays(nums []int, minK int, maxK int) int64 {
    var res int64 
    iMin, iMax, iBefore := -1, -1, -1
    for i, k := range nums {
        if k == minK {
            iMin = i
        }
        if k == maxK {
            iMax = i
        }
        if k < minK || k > maxK {
            iMin, iMax, iBefore = -1, -1, i
        }
        if iMin >= 0 && iMax >= 0 {
            res += int64(min(iMin, iMax) - iBefore)
        }
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}