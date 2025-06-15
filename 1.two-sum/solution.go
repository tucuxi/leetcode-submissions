func twoSum(nums []int, target int) []int {
    h := map[int]int{}
    for i, n := range nums {
        if j, ok := h[target - n]; ok {
            return []int{j, i}
        }
        h[n] = i
    }
    return nil
}