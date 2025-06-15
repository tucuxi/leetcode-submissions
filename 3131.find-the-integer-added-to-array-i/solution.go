func addedInteger(nums1 []int, nums2 []int) int {
    return minElement(nums2) - minElement(nums1)
}

func minElement(nums []int) int {
    res := math.MaxInt
    for _, n := range nums {
        if n < res {
            res = n
        }
    }
    return res
}