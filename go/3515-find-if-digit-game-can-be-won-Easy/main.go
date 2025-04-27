func canAliceWin(nums []int) bool {
    b := 0
    for _, n := range nums {
        if n < 10 {
            b += n
        } else {
            b -= n
        }
    }
    return b != 0
}