func countSpecialIntegers(nums []int) int {
    h := make([]int, 101)
    p := 0

    for _, x := range nums {
        if x != p {
            h[x]++
            p = x
        }
    }

    res := 0
    for x := 1; x < len(h); x++ {
        if h[x] == 1 {
            res++
        }
    }

    return res
}