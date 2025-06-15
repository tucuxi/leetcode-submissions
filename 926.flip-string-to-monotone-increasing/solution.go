func minFlipsMonoIncr(s string) int {
    flips, ones := 0, 0
    for _, r := range s {
        if r == '0' {
            flips++
        } else {
            ones++
        }
        if ones < flips {
            flips = ones
        }
    }
    return flips
}