func countRotations(s string, k int) int {
    n := len(s)
    res := 0

    for i := range n {
        c := 0
        for j := range n-1 {
            if s[(i+j) % n] == s[(i+j+1) % n] {
                c++
            }
        }
        if c == k {
            res++
        }
    }

    return res
}