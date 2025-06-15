func countGood(nums []int, k int) int64 {
    res := int64(0)
    count := make(map[int]int)
    same := 0
    j := -1
    for i := range nums {
        for same < k && j + 1 < len(nums) {
            j++
            same += count[nums[j]]
            count[nums[j]]++
        }
        if same >= k {
            res += int64(len(nums) - j)
        }
        count[nums[i]]--
        same -= count[nums[i]]
    }
    return res
}