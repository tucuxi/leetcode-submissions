func minBitwiseArray(nums []int) []int {
    var c [1001]int

    for i := range c {
        c[i] = -1
    }
    for i := 1; i < len(c); i++ {
        x := i | (i + 1)
        if x < len(c) && c[x] == -1 {
            c[x] = i
        }
    }

    var ans []int

    for _, n := range nums {
        ans = append(ans, c[n])
    }

    return ans
}
