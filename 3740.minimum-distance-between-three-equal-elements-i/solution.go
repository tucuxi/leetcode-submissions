func minimumDistance(nums []int) int {
    res := math.MaxInt
    for i := range nums {
        for j := i+1; j < len(nums); j++ {
            if nums[j] == nums[i] {
                for k := j+1; k < len(nums); k++ {
                    if nums[k] == nums[i] {
                        res = min(res, 2*(k-i))
                    }
                }
            }
        }
    }
    if res == math.MaxInt {
        return -1
    }
    return res
}