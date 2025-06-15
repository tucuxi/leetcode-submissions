func findScore(nums []int) int64 {
    bls := make(map[int]bool)
    idx := make([]int, len(nums))
    for i := range idx {
        idx[i] = i
    }
    sort.SliceStable(idx, func(i, j int) bool { return nums[idx[i]] < nums[idx[j]] })
    var score int64
    for i := range idx {
        index := idx[i]
        if !bls[index] {
            score += int64(nums[index])
            if index > 0 {
                bls[index - 1] = true
            }
            if index + 1 < len(nums) {
                bls[index + 1] = true
            }
        } 
    }
    return score
}