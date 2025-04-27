func minNumber(nums1 []int, nums2 []int) int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    i, j := 0, 0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] == nums2[j] {
            return nums1[i]
        }
        if nums1[i] < nums2[j] {
            i++
        } else {
            j++
        }
    }
    if nums1[0] < nums2[0] {
        return 10 * nums1[0] + nums2[0]
    }
    return 10 * nums2[0] + nums1[0]
}