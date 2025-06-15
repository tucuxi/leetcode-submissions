func increasingTriplet(nums []int) bool {
    p, q := math.MaxInt, math.MaxInt
    for _, n := range nums {
        if n <= p {
            p = n
        } else if n <= q {
            q = n
        } else {
            return true
        }
    }
    return false
}