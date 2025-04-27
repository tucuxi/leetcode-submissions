func findDiagonalOrder(nums [][]int) []int {
    m := make(map[int][]int)
    maxKey := -1
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums[i]); j++ {
            key := i + j
            if key > maxKey {
                maxKey = key
            }
            m[key] = append([]int{ nums[i][j] }, m[key]...)
        }
    }
    res := []int{}
    for i := 0; i <= maxKey; i++ {
        res = append(res, m[i]...)
    }
    return res
}