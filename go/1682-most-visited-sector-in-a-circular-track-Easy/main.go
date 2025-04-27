func mostVisited(n int, rounds []int) []int {
    count := make([]int, n)
    for i := 1; i < len(rounds); i++ {
        m := (rounds[i] - rounds[i - 1] + n) % n
        k := rounds[i - 1] - 1
        for j := 0; j < m; j++ {
            count[k]++
            k++
            if (k == n) {
                k = 0
            }
        }
    }
    count[rounds[len(rounds) - 1] - 1]++
    max := 0
    for _, v := range count {
        if v > max {
            max = v
        }
    }
    var res []int
    for i, v := range count {
        if v == max {
            res = append(res, i + 1)
        }
    }
    return res
}