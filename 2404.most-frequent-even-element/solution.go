func mostFrequentEven(nums []int) int {
    sort.Ints(nums)
    res, prev := -1, -1
    frq, max := 0, 0
    for _, n := range nums {
        if n == prev {
            frq++
        } else {
            frq = 1
            prev = n
        }
        if n % 2 == 0 && frq > max {
            max = frq
            res = n
        }
    }
    return res
}