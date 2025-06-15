func accountsMerge(accounts [][]string) [][]string {
    u := makeUF(len(accounts))
    h := make(map[string]int)
    for i, a := range accounts {
        for j := 1; j < len(a); j++ {
            if other, exists := h[a[j]]; exists {
                u.union(i, other)
            } else {
                h[a[j]] = i
            }
        }
    }
    m := make(map[int][]string)
    for i, a := range accounts {
        b := u.find(i)
        for j := 1; j < len(a); j++ {
            m[b] = append(m[b], a[j])
        }
    }
    res := make([][]string, 0)
    for i, arr := range m {
        sort.Strings(arr)
        k := 1
        for j := 1; j < len(arr); j++ {
            if arr[j] != arr[k - 1] {
                arr[k] = arr[j]
                k++
            }
        }
        res = append(res, append([]string{accounts[i][0]}, arr[:k]...))
    }
    return res
}

type uf struct {
    parent []int
    size []int
}

func makeUF(n int) uf {
    parent := make([]int, n)
    size := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
        size[i] = 1
    }
    return uf{parent, size}
}

func (u *uf) union(x, y int) {
    x = u.find(x)
    y = u.find(y)
    if x == y {
        return
    }
    if u.size[x] > u.size[y] {
        u.parent[y] = x
        u.size[x] += u.size[y]
    } else {
        u.parent[x] = y
        u.size[y] += u.size[x]
    }
}

func (u *uf) find(x int) int {
    for x != u.parent[x] {
        x, u.parent[x] = u.parent[x], u.parent[u.parent[x]]
    }
    return x
}

func (u *uf) group(p int) []int {
    res := make([]int, 0)
    for i := range u.parent {
        if u.parent[i] == p {
            res = append(res, i)
        }
    }
    return res
}