func canMakeArithmeticProgression(arr []int) bool {
    min, max := math.MaxInt, math.MinInt
    for _, a := range arr {
        if a < min {
            min = a
        }
        if a > max {
            max = a
        }
    }
    r := max - min
    if r == 0 {
        return true
    }
    if r % (len(arr) - 1) != 0 {
        return false
    }
    j := r / (len(arr) - 1)
    set := make(map[int]bool)
    for _, a := range arr {
        set[(a - min) / j] = true
    }
    return len(set) == len(arr)
}