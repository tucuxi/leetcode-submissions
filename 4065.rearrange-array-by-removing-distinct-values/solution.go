func rearrangeArray(nums []int) []int {
    var h [101]int

    for _, num := range nums {
        h[num]++
    }

    ans := make([]int, 0, len(nums))

    for len(ans) < len(nums) {
        for num := 1; num < len(h); num++ {
            if h[num] > 0 {
                ans = append(ans, num)
                h[num]--
            }
        }
    }

    return ans
}