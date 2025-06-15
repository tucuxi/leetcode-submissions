func findNumberOfLIS(nums []int) int {
    m := make([]int, len(nums))
    c := make([]int, len(nums))
    mmax := 0
    cmax := 0
    for i := range nums {
        m[i] = 1
        c[i] = 1
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                if m[j] + 1 > m[i] {
                    m[i] = m[j] + 1
                    c[i] = c[j]
                } else if m[j] + 1 == m[i] {
                    c[i] += c[j]
                }
            }   
        }
        if m[i] > mmax {
            mmax = m[i]
            cmax = c[i]
        } else if m[i] == mmax {
            cmax += c[i]
        }
    }
    return cmax
}