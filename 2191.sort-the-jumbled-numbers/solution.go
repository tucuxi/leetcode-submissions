func sortJumbled(mapping []int, nums []int) []int {

    mapInt := func(n int) int {
        if n == 0 {
            return mapping[0]
        }
        res := 0
        f := 1
        for n > 0 {
            d := n % 10
            res += mapping[d] * f
            f *= 10
            n /= 10
        }
        return res
    }    

    mappedNums := make([]int, len(nums))
    index := make([]int, len(nums))
    for i, n := range nums {
        mappedNums[i] = mapInt(n)
        index[i] = i
    }

    sort.SliceStable(index, func(i, j int) bool {
        return mappedNums[index[i]] < mappedNums[index[j]]
    })

    res := make([]int, len(nums))
    for i := range res {
        res[i] = nums[index[i]]
    }

    return res
}