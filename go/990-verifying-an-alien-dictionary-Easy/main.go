func isAlienSorted(words []string, order string) bool {
    m := map[byte]int{}
    for i := range order {
        m[order[i]] = i
    }
    for i := 1; i < len(words); i++ {
        k := len(words[i-1])
        if len(words[i]) < k {
            k = len(words[i])
        }
        j := 0
        for ; j < k && words[i-1][j] == words[i][j]; j++ {
        }
        if j < k && m[words[i-1][j]] > m[words[i][j]] {
            return false
        }
        if j == k && j < len(words[i-1]) {
            return false
        } 
    }
    return true
}