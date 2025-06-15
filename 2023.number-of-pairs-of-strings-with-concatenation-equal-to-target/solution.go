func numOfPairs(nums []string, target string) int {
    res := 0
    for i := range nums {
        if len(nums[i]) < len(target) && nums[i] == target[:len(nums[i])] {
            for j := range nums {
                if i != j && nums[j] == target[len(nums[i]):] {
                    res++
                }
            }
        }
    }
    return res
}