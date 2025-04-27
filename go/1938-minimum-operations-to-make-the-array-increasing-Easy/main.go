func minOperations(nums []int) int {
    c := 0
    p := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] <= p {
            c += p - nums[i] + 1
            p++
        } else {
            p = nums[i]
        }
    }
    return c
}