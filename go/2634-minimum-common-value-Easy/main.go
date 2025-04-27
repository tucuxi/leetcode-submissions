func getCommon(nums1 []int, nums2 []int) int {
    if len(nums1) < len(nums2) {
        nums1, nums2 = nums2, nums1
    }
    for _, n := range nums1 {
        _, found := slices.BinarySearch(nums2, n)
        if found {
            return n
        }
    }
    return -1
}