func longestSubsequence(nums []int) int {
    x := 0
    z := true

    for _, num := range nums {
        x ^= num
        z = z && num == 0
    }

    if z {
        return 0
    }
    if x == 0 {
        return len(nums) - 1
    }
    return len(nums)
}