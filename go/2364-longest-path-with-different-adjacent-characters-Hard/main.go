func longestPath(parent []int, s string) int {
    children := make([][]int, len(parent))
    for i, p := range parent {
        if p != -1 {
            children[p] = append(children[p], i)
        }
    }
    maxPath := 0

    var dfs func(int) int
    dfs = func(p int) int {
        l1, l2 := 0, 0
        for _, c := range children[p] {
            l := dfs(c)
            if s[c] != s[p] {
                if l >= l1 {
                    l1, l2 = l, l1
                } else if l > l2 {
                    l2 = l
                }
            }
        }
        path := l1 + l2 + 1
        if path > maxPath {
            maxPath = path
        }
        return l1 + 1
    }

    dfs(0)
    return maxPath
}