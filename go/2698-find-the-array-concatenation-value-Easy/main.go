func findTheArrayConcVal(nums []int) int64 {
    var res int64
    for i, j := 0, len(nums) - 1; i <= j; i, j = i+1, j-1 {
        a := nums[i]
        if i < j {
            for b := nums[j]; b > 0; b /= 10 {
                a *= 10
            }
            a += nums[j]
        }
        res += int64(a)
    }
    return res
}