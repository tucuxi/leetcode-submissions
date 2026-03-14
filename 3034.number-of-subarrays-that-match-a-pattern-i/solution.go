func countMatchingSubarrays(nums []int, pattern []int) int {
    res := 0

    outer: for i := range len(nums) - len(pattern) {
        for k := range pattern {
            if sgn(nums[i + k + 1] - nums[i + k]) != pattern[k] {
                continue outer
            }
        }
        res++
    }
    
    return res
}

func sgn(a int) int {
    if a < 0 {
        return -1
    }
    if a > 0 {
        return 1
    }
    return 0
}