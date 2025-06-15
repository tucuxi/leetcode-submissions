func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
    u := newUF(n)
    for x := 0; x < n; x++ {
        for _, c := range []int{leftChild[x], rightChild[x]} {
            if c != -1 && (u.find(c) != c || !u.union(x, c)) {
                return false
            }
        }
    }
    return u.groups() == 1
}

type uf struct {
    parent []int
}

func newUF(n int) *uf {
    parent := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
    }
    return &uf{parent}
}

func (u *uf) union(x, y int) bool {
    x = u.find(x)
    y = u.find(y)
    if x == y {
        return false
    }
    u.parent[y] = x
    return true
}

func (u *uf) find(x int) int {
    for x != u.parent[x] {
        x, u.parent[x] = u.parent[x], u.parent[u.parent[x]]
    }
    return x
}

func (u *uf) groups() int {
    res := 0
    for x := range u.parent {
        if x == u.parent[x] {
            res++
        }
    }
    return res
}