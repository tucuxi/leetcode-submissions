func xorAllNums(nums1 []int, nums2 []int) int {
    ans := 0
    if len(nums2) % 2 != 0 {
        for _, n := range nums1 {
            ans ^= n
        }
    }
    if len(nums1) % 2 != 0 {
        for _, n := range nums2 {
            ans ^= n
        }
    }
    return ans
}