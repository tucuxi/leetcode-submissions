func jump(nums []int) int {
    far := 0
    end := 0
    count := 0
    for i := 0; i < len(nums) - 1; i++ {
        if i + nums[i] > far {
            far = i + nums[i]
        }
        if i == end {
            count++
            end = far
        }
    }
    return count
}