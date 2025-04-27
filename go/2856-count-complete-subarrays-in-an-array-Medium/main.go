func countCompleteSubarrays(nums []int) int {
    distinct := countDistinct(nums)
    hs := make(map[int]int)
    r := 0
    res := 0

    for l := range nums {
        if l > 0 {
            p := nums[l - 1]
            hs[p]--
            if hs[p] == 0 {
                delete(hs, p)
            }
        }
        for r < len(nums) && len(hs) < distinct {
            hs[nums[r]]++
            r++
        }
        if len(hs) == distinct {
            res += len(nums) - r + 1
        }
    }
    return res
}

func countDistinct(nums []int) int {
    h := make(map[int]bool)
    for _, num := range nums {
        h[num] = true
    }
    return len(h)
}