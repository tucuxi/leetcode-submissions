func returnToBoundaryCount(nums []int) int {
    c := 0
    p := 0
    for _, n := range nums {
        p += n
        if p == 0 {
            c++
        }
    }
    return c
}