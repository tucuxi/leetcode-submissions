func findNumbers(nums []int) int {
    res := 0
    for _, num := range nums {
        even := true
        for num > 0 {
            even = !even
            num /= 10
        }
        if even {
            res++
        }
    }
    return res
}