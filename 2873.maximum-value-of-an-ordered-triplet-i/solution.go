func maximumTripletValue(nums []int) int64 {
    var ans int64
    for i := range nums {
        for j := i + 1; j < len(nums); j++ {
            h := int64(nums[i] - nums[j])
            for k := j + 1; k < len(nums); k++ {
                if v := h * int64(nums[k]); v > ans {
                    ans = v
                }
            }
        }
    }
    return ans
}