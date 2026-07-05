func isMiddleElementUnique(nums []int) bool {
    m := nums[len(nums) / 2]
    c := 0
    for _, num := range nums {
        if num == m {
            c++
        }
    }
    return c == 1
}