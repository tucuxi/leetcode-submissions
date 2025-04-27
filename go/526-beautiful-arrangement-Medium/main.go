func countArrangement(n int) int {
    var count int
    var perm func(int, []int, []bool)

    perm = func(index int, path []int, used []bool) {
        if index > n {
            count++
            return
        }
        for i := 1; i <= n; i++ {
            if !used[i] && (index % i == 0 || i % index == 0) {
                path[index] = i
                used[i] = true
                perm(index + 1, path, used)
                used[i] = false
            }
        }
    }

    perm(1, make([]int, n + 1), make([]bool, n + 1))
    return count
}