func leastInterval(tasks []byte, n int) int {
    var f [26]int
    m := 0
    for _, t := range tasks {
        index := t - 'A'
        f[index]++
        m = max(m, f[index])
    }
    p := 0
    for _, c := range f {
        if c == m {
            p++
        }
    }
    return max(len(tasks), (n + 1) * (m - 1) + p)
}