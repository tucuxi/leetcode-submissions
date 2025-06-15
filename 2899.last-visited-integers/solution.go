func lastVisitedIntegers(words []string) []int {
    var res []int
    nums := []int{-1}
    k := 0
    for _, w := range words {
        n, err := strconv.Atoi(w)
        if err != nil {
            if k < len(nums) {
                k++
            }
            res = append(res, nums[len(nums) - k])
        } else {
            k = 0
            nums = append(nums, n)
        }
    }
    return res
}