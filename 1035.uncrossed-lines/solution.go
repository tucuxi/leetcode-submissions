func maxUncrossedLines(nums1 []int, nums2 []int) int {
    ul := make([][]int, len(nums1) + 1)
    for i := range ul {
        ul[i] = make([]int, len(nums2) + 1)
    }
    for i := range nums1 {
        for j := range nums2 {
            if nums1[i] == nums2[j] {
                ul[i + 1][j + 1] = ul[i][j] + 1
            } else {
                ul[i + 1][j + 1] = max(ul[i + 1][j], ul[i][j + 1])
            }
        }
    }
    return ul[len(nums1)][len(nums2)]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}