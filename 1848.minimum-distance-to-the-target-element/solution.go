func getMinDistance(nums []int, target int, start int) int {
    for offset := 0; ; offset++ {
        if i := start+offset; i < len(nums) && nums[i] == target {
            return offset
        }
        if i := start-offset; i >= 0 && nums[i] == target {
            return offset
        }
    } 
}