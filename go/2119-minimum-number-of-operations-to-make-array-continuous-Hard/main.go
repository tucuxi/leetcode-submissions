func minOperations(nums []int) int {
    h := make(map[int]bool)
    for _, n := range nums {
        h[n] = true
    }
    usn := make([]int, 0, len(h))
    for n := range h {
        usn = append(usn, n)
    }
    sort.Ints(usn)
    res := len(nums)
    j := 0
    for i, n := range usn {
        for j < len(usn) && usn[j] < n + len(nums) {
            j++
        }
        if k := len(nums) - (j - i); k < res {
            res = k
        }
    }
    return res
}