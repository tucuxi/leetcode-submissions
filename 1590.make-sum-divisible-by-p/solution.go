func minSubarray(nums []int, p int) int {
    remainder := 0
    for _, n := range nums {
        remainder = (remainder + n) % p
    }
    if remainder == 0 {
        return 0
    }
    modMap := map[int]int{0: -1}
    currentSum := 0
    minLen := len(nums)
    for i, n := range nums {
        currentSum = (currentSum + n) % p
        needed := (currentSum - remainder + p) % p
        if j, exists := modMap[needed]; exists {
            minLen = min(minLen, i - j)
        }
        modMap[currentSum] = i
    }
    if minLen == len(nums) {
        return -1
    }
    return minLen
}