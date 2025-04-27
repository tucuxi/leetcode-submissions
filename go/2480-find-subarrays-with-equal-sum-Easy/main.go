func findSubarrays(nums []int) bool {
    subs := map[int]bool{}
    for i := 1; i < len(nums); i++ {
        s := nums[i-1] + nums[i]
        if subs[s] {
            return true
        }
        subs[s] = true
    }
    return false
}