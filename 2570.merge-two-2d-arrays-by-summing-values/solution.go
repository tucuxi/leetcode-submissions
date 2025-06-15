func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    res := make([][]int, 0, len(nums1) + len(nums2))
    for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
        if j == len(nums2) || i < len(nums1) && nums1[i][0] < nums2[j][0] {
            res = append(res, nums1[i])
            i++
        } else if i == len(nums1) || nums1[i][0] > nums2[j][0] {
            res = append(res, nums2[j])
            j++
        } else {
            res = append(res, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
            i++
            j++
        }
    }
    return res
}