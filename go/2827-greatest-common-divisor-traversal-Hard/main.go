var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313}

func canTraverseAllPairs(nums []int) bool {
    if len(nums) == 1 {
        return true
    }
    u := newUnionFind(100001)
    for _, n := range nums {
        if n == 1 {
            return false
        }
        k := n
        for _, p := range primes {
            if k % p == 0 {
                u.union(n, p)
                for k % p == 0 {
                    k /= p
                }
            }
        }
        if k > 1 {
            u.union(n, k)
        }
    }
    root := u.find(nums[0])
    for _, n := range nums {
        if u.find(n) != root {
            return false
        }
    }
    return true
}


type unionFind struct {
    parent []int
    size []int
}

func newUnionFind(n int) *unionFind {
    p := make([]int, n)
    s := make([]int, n)
    for i := range p {
        p[i] = i
        s[i] = 1
    }
    return &unionFind{p, s}
}

func (u *unionFind) union(x, y int) {
    x = u.find(x)
    y = u.find(y)
    if x == y {
        return
    }
    if u.size[x] < u.size[y] {
        u.parent[x] = y
        u.size[y] += u.size[x]
    } else {
        u.parent[y] = x
        u.size[x] += u.size[y]
    }
}

func (u *unionFind) find(x int) int {
    for x != u.parent[x] {
        x, u.parent[x] = u.parent[x], u.parent[u.parent[x]]
    }
    return x
}