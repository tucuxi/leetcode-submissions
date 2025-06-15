func combine(n int, k int) [][]int {
    combinations := [][]int{}
    prefix := make([]int, 0, k)
    
    var rec func(int, int)
    rec = func(remaining int, next int) {
        if remaining == 0 {
            combination := make([]int, k)
            copy(combination, prefix)
            combinations = append(combinations, combination)
            return
        }
        for i := next; i <= n; i++ {
            prefix = append(prefix, i)
            rec(remaining - 1, i + 1)
            prefix = prefix[:len(prefix) - 1]
        }
    }

    rec(k, 1)
    return combinations
}