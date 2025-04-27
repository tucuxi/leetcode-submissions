func divideArray(nums []int) bool {
    var h [501]bool
    c := 0
    for _, num := range nums {
        if h[num] {
            c--
            h[num] = false
        } else {
            c++
            h[num] = true
        }
    }
    return c == 0
}