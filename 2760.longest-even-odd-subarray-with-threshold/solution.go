func longestAlternatingSubarray(nums []int, threshold int) int {
    res := 0
    l := 0
    for l < len(nums) {
        n := nums[l]
        if n > threshold || n % 2 != 0 { 
            l++
            continue
        }
        r := l + 1
        for r < len(nums) && nums[r] <= threshold && nums[r] % 2 != n % 2 {
            n = nums[r]
            r++
        }
        if r - l > res {
            res = r - l
        }
        l = r
    }
    return res
}