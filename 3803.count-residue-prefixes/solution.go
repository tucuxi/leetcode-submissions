func residuePrefixes(s string) int {
    visited := make([]bool, 26)
    distinct := 0
    res := 0

    for i, ch := range s {
        if index := ch - 'a'; !visited[index] {
            visited[index] = true
            distinct++
        }
        if distinct == (i + 1) % 3 {
            res++
        }
    }

    return res
}