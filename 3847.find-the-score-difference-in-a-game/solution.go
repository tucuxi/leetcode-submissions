func scoreDifference(nums []int) int {
    p := []int{0, 0}
    active := 0
    for i := range nums {
        if nums[i] % 2 != 0 {
            active = 1 - active
        }
        if (i + 1) % 6 == 0 {
            active = 1 - active
        }
        p[active] += nums[i]
    }
    return p[0] - p[1]
}