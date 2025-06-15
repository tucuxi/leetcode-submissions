func pathInZigZagTree(label int) []int {
    var p []int
    for n := label / 2; n > 0; n /= 2 {
        p = append(p, n)
    }
    slices.Reverse(p)
    res := make([]int, 0, len(p) + 1)
    reverse := len(p) % 2 != 0
    mask := 0
    for _, n := range p {
        if reverse {
            n ^= mask
        }
        res = append(res, n)
        reverse = !reverse
        mask = (mask << 1) | 1
    }
    res = append(res, label)
    return res    
}