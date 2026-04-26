func maxNonOverlapping(nums []int, target int) int {
    res := 0
    v := map[int]bool{0: true}
    s := 0

    for _, n := range nums {
        s += n
        r := s-target
        if v[r] {
            res++
            s = 0
            v = make(map[int]bool)
        }
        v[s] = true
    }

    return res
}