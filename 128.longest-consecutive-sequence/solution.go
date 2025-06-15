func longestConsecutive(nums []int) int {
    set := map[int]bool{}
    for _, n := range nums {
        set[n] = true
    }
    max := 0
    for _, n := range nums {
        if set[n] {
            l := n
            for set[l] {
                delete(set, l)
                l--
            }
            r := n + 1
            for set[r] {
                delete(set, r)
                r++
            }
            if len := r - l - 1; len > max {
                max = len
            }
        }
    }
    return max
}