func unequalTriplets(nums []int) int {
    res := 0
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] == nums[j] {
                continue
            }
            for k := j + 1; k < len(nums); k++ {
                if nums[i] != nums[k] && nums[j] != nums[k] {
                    res++
                }
            }
        }
    }
    return res
}