func findKDistantIndices(nums []int, key int, k int) []int {
    h := make(map[int]bool)
    for i := range nums {
        if nums[i] == key {
            for j := 0; j <= k; j++ {
                if i - j >= 0 {
                    h[i - j] = true
                }
                if i + j < len(nums) {
                    h[i + j] = true
                }
            }
        }
    }
    var res []int
    for i := range nums {
        if h[i] {
            res = append(res, i)
        }
    }
    return res
}