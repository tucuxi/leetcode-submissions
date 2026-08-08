func maxPairStrength(nums []int) int64 {
    var res int64
    for i := range nums {
        for j := i+1; j < len(nums); j++ {
            gcd := gcd(nums[i], nums[j])
            strength := int64(nums[i] / gcd) * int64(nums[j] / gcd)
            res = max(res, strength)
        }
    }
    return res
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}