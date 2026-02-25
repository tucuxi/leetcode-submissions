func firstUniqueFreq(nums []int) int {
    h := make(map[int]int)
    for _, num := range nums {
        h[num]++
    }

    f := make(map[int]int)
    for _, count := range h {
        f[count]++
    }

    for _, num := range nums {
        if f[h[num]] == 1 {
            return num
        }
    }

    return -1
}