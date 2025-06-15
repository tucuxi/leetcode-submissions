func minSubArrayLen(target int, nums []int) int {
    min := math.MaxInt
    sum := 0
    p, q := 0, 0
    for q < len(nums) {
        if nums[q] >= target {
            return 1
        }
        if nums[q] + sum >= target {
            if q - p + 1 < min {
                min = q - p + 1
            }
            sum -= nums[p]
            p++
        } else {
            sum += nums[q]
            q++
        }
    }
    if min == math.MaxInt {
        return 0
    }
    return min
}