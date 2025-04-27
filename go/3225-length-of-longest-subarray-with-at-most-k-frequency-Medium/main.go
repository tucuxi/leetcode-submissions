func maxSubarrayLength(nums []int, k int) int {
    h := make(map[int]int)
    run, max := 0, 0
    for i, n := range nums {
        if h[n] == k {
            j := i - run
            for nums[j] != n {
                h[nums[j]]--
                j++
            }    
            run = i - j
        } else {
            h[n]++
            run++
            if run > max {
                max = run
            }
        }
    }
    return max
}