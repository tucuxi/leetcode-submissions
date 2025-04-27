func beautifulSubsets(nums []int, k int) int {
    res := 0

    var rec func(int, []int)
    rec = func(index int, subset []int) {
        if index == len(nums) {
            res++
            return
        }
        seen := false
        for _, num := range subset {
            seen = seen || nums[index] - num == k || num - nums[index] == k 
        }
        if !seen {
            rec(index+1, append(subset, nums[index]))
        }
        rec(index+1, subset)
    }

    rec(0, make([]int, 0, len(nums)))
    return res-1
}