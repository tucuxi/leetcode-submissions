func findLength(nums1 []int, nums2 []int) int {
    n1, n2 := len(nums1), len(nums2)
    dp := make([][]int, n1 + 1)
    for i := range dp {
        dp[i] = make([]int, n2 + 1)
    }
    res := 0
    for i := 0; i < n1; i++ {
        for j := 0; j < n2; j++ {
            if nums1[i] == nums2[j] {
                k := dp[i][j] + 1
                dp[i + 1][j + 1] = k 
                if k > res {
                    res = k 
                }
            }
        }
    }
    return res
}