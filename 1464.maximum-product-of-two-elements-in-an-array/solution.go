func maxProduct(nums []int) int {
    m1, m2 := 0, 0
    for _, num := range nums {
        if num > m1 {
            m1, m2 = num, m1
        } else if num > m2 {
            m2 = num
        }
    }
    return (m1-1) * (m2-1)
}