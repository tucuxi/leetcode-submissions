func numSubarraysWithSum(nums []int, goal int) int {
    p := make(map[int]int)
    p[0] = 1
    s := 0
    res := 0
    for _, n := range nums {
        s += n
        res += p[s - goal]
        p[s]++ 
    }
    return res
}