func minimumRounds(tasks []int) int {
    h := map[int]int{}
    for _, difficulty := range tasks {
        h[difficulty]++
    }
    res := 0
    for _, count := range h {
        if count == 1 {
            return -1
        }
        res += (count + 2) / 3
    }
    return res
}