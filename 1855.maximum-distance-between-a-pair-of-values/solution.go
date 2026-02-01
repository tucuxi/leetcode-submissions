func maxDistance(nums1 []int, nums2 []int) int {
    i, j, res := 0, 0, 0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] <= nums2[j] {
            res = max(res, j - i)
            j++
        } else {
            i++
        }
    }
    return res
}