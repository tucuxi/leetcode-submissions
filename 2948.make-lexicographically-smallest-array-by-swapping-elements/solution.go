func lexicographicallySmallestArray(nums []int, limit int) []int {
    index := func() []int {
        res := make([]int, len(nums))
        for i := range res {
            res[i] = i
        }
        slices.SortFunc(res, func(a, b int) int {
            return nums[a] - nums[b]
        })
        return res
    }()

    return func() []int {
        res := make([]int, len(nums))
        i := 0
        for i < len(nums) {
            start := i
            for i++; i < len(nums) && nums[index[i]] <= nums[index[i - 1]] + limit; i++ {}
            partition := slices.Clone(index[start:i])
            slices.Sort(partition)
            for j := start; j < i; j++ {
                res[partition[j - start]] = nums[index[j]]
            }
        }
        return res
    }()
}