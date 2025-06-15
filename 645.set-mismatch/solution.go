func findErrorNums(nums []int) []int {
    h := make([]int, len(nums) + 1)
    for _, n := range nums {
        h[n]++
    }
    res := make([]int, 2)
    for i := 1; i <= len(nums); i++ {
        switch h[i] {
            case 0: res[1] = i
            case 2: res[0] = i
        } 
    }
    return res
}