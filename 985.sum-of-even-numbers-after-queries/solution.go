func sumEvenAfterQueries(nums []int, queries [][]int) []int {
    even := 0
    for _, n := range nums {
        if n % 2 == 0 {
            even += n
        }
    }
    res := []int{}
    for _, q := range queries {
        v, i := q[0], q[1]
        if nums[i] % 2 == 0 {
            even -= nums[i]
        }
        nums[i] += v
        if nums[i] % 2 == 0 {
            even += nums[i]
        }
        res = append(res, even)
    }
    return res
}