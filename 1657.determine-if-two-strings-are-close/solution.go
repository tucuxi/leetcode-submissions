const n = 26

func closeStrings(word1 string, word2 string) bool {
    f1, f2 := make([]int, n), make([]int, n)
    s1, s2 := [n]bool{}, [n]bool{}
    for i := range word1 {
        c := word1[i] - 'a'
        f1[c]++
        s1[c] = true
    }
    for i := range word2 {
        c := word2[i] - 'a'
        f2[c]++
        s2[c] = true
    }
    if s1 != s2 {
        return false
    }
    sort.Ints(f1)
    sort.Ints(f2)
    for i := 0; i < n; i++ {
        if f1[i] != f2[i] {
            return false
        }
    }
    return true
}