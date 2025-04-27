func countElements(nums []int) int {
    min, max := 100001, -100001
    for _, n := range nums {
        if (n < min) {
            min = n
        }
        if (n > max) {
            max = n
        }
    }
    count := 0
    for _, n := range nums {
        if (n > min && n < max) {
            count++
        }
    }
    return count
}