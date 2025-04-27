func longestNiceSubarray(nums []int) int {
    res := 0
    l := 0
    m := 0
    for r := range nums {
        for nums[r] & m != 0 {
            m ^= nums[l]
            l++
        }
        res = max(res, r - l + 1)
        m |= nums[r]
    }
    return res
}