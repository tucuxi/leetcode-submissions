func constructDistancedSequence(n int) []int {
    var (
        constructRecursivelyWithBacktracking func(int) bool
        res = make([]int, 2*n - 1)
        visited = make([]bool, n + 1) 
    )

    constructRecursivelyWithBacktracking = func(index int) bool {
        if index == len(res) {
            return true
        }
        if res[index] != 0 {
            return constructRecursivelyWithBacktracking(index + 1)
        }
        for i := n; i > 0; i-- {
            if visited[i] {
                continue
            }
            visited[i] = true
            res[index] = i
            if i == 1 {
                if constructRecursivelyWithBacktracking(index + 1) {
                    return true
                }
            } else {
                if index + i < len(res) && res[index + i] == 0 {
                    res[index + i] = i
                    if constructRecursivelyWithBacktracking(index + 1) {
                        return true
                    }
                    res[index + i] = 0 
                }
            }
            res[index] = 0
            visited[i] = false
        }
        return false
    }

    ok := constructRecursivelyWithBacktracking(0)
    if !ok {
        panic("no solution found")
    }
    return res
}