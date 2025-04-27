func mostFrequent(nums []int, key int) int {
    target := map[int]int{}
    for i := 0; i <= len(nums) - 2; i++ {
        if nums[i] == key {
            target[nums[i+1]]++
        }
    }
    maxCount := 0
    res := 0
    for n, count := range target {
        if count > maxCount {
            maxCount = count
            res = n
        }
    }
    return res
}