func maxFrequencyElements(nums []int) int {
    var f [101]int
    var max, res int
    
    for _, n := range nums {
        f[n]++
        if f[n] == max {
            res += max
        } else if f[n] > max { 
            max = f[n]
            res = max
        }
    } 
    
    return res
}