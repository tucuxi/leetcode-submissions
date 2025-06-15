func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
    edgesWithIndex := func() [][4]int {
        res := make([][4]int, len(edges))
        for i := range edges {
            res[i][0] = edges[i][0]
            res[i][1] = edges[i][1]
            res[i][2] = edges[i][2]
            res[i][3] = i
        }
        sort.Slice(res, func(i, j int) bool { return res[i][2] < res[j][2] })
        return res
    }()
    
    mstWeight := func() int {
        weight := 0
        u := newUf(n)
        for _, e := range edgesWithIndex {
            if u.union(e[0], e[1]) {
                weight += e[2]
            }
        }
        return weight
    }()
        
    res := make([][]int, 2)
    
    for i := range edgesWithIndex {
        ignoreWeight, connected := func() (int, bool) {
            uf := newUf(n)
            weight := 0
            for j := range edgesWithIndex {
                if j != i && uf.union(edgesWithIndex[j][0], edgesWithIndex[j][1]) {
                    weight += edgesWithIndex[j][2]
                }
            }
            return weight, uf.connected()
        }()
        
        if !connected || ignoreWeight > mstWeight {
            res[0] = append(res[0], edgesWithIndex[i][3])
            continue
        }
        
        forceWeight := func() int {
            uf := newUf(n)
            uf.union(edgesWithIndex[i][0], edgesWithIndex[i][1])
            weight := edgesWithIndex[i][2]
            for j := range edgesWithIndex {
                if j != i && uf.union(edgesWithIndex[j][0], edgesWithIndex[j][1]) {
                    weight += edgesWithIndex[j][2]
                }
            }
            return weight
        }()
        
        if forceWeight == mstWeight {
            res[1] = append(res[1], edgesWithIndex[i][3])
        }
    }
    
    return res
}

type uf struct {
    parent []int
    size []int
    maxSize int
}

func newUf(n int) *uf {
    u := new(uf)
    u.parent = make([]int, n)
    u.size = make([]int, n)
    u.maxSize = 1
    for i := range u.parent {
        u.parent[i] = i
        u.size[i] = 1
    }
    return u
}

func (u *uf) union(p, q int) bool {
    p = u.find(p)
    q = u.find(q)
    if p == q {
        return false
    }
    if u.size[p] < u.size[q] {
        p, q = q, p
    }
    u.parent[q] = p
    u.size[p] += u.size[q]
    if u.size[p] > u.maxSize {
        u.maxSize = u.size[p]
    }
    return true
}

func (u *uf) find(p int) int {
    for u.parent[p] != p {
        p, u.parent[p] = u.parent[p], u.parent[u.parent[p]]
    }
    return p
}

func (u *uf) connected() bool {
    return u.maxSize == len(u.parent)
}