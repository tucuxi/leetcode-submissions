func minCapability(nums []int, k int) int {
    maxCash := nums[0]

    for _, cash := range nums[1:] {
        maxCash = max(maxCash, cash)
    }

    return sort.Search(maxCash + 1, func(cash int) bool {
        c := 0
        for i := 0; i < len(nums); i++ {
            if nums[i] <= cash {
                c++
                i++
            }
        }
        return c >= k
    })
}