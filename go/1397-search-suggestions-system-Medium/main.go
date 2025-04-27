func suggestedProducts(products []string, searchWord string) [][]string {
    sort.Strings(products)
    m := len(searchWord)
    res := make([][]string, m)
    p := products
    for i := 0; i < m; i++ {
        q := []string{}
        for j := 0; j < len(p); j++ {
            if len(p[j]) > i && p[j][i] == searchWord[i] {
                q = append(q, p[j])
            }
        }
        k := min(len(q), 3)
        res[i] = q[:k]
        p = q
    }
    return res
}

func min(a, b int) int {
    if (a < b) {
        return a
    } else {
        return b
    }
}