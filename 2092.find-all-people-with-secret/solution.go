func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
    ds := newDisjointSet(n)
    ds.union(0, firstPerson)
    sort.Slice(meetings, func(i, j int) bool { return meetings[i][2] < meetings[j][2] })
    
    for i := 0; i < len(meetings); {
        j := i
        for ; j < len(meetings) && meetings[j][2] == meetings[i][2]; j++ {
            ds.union(meetings[j][0], meetings[j][1])
        }
        for ; i < j; i++ {
            if !ds.connected(meetings[i][0], 0) {
                ds.reset(meetings[i][0])
                ds.reset(meetings[i][1])
            }
        }
    }
    
    var res []int
    for x := 0; x < n; x++ {
        if ds.connected(x, 0) {
            res = append(res, x)
        }
    }
    return res
}

type disjointSet struct {
    parent []int
    rank []int
}

func newDisjointSet(n int) *disjointSet {
    parent := make([]int, n)
    for i := range parent {
        parent[i] = i
    }
    return &disjointSet{
        parent: parent,
        rank: make([]int, n),
    }
}

func (d *disjointSet) union(x, y int) {
    x = d.find(x)
    y = d.find(y)
    if x == y {
        return
    }
    if d.rank[x] < d.rank[y] {
        d.parent[x] = y
    } else if d.rank[x] > d.rank[y] {
        d.parent[y] = x
    } else {
        d.parent[y] = x
        d.rank[x]++
    }
}

func (d *disjointSet) find(x int) int {
    for x != d.parent[x] {
        x, d.parent[x] = d.parent[x], d.parent[d.parent[x]]
    }
    return x 
}

func (d *disjointSet) connected(x, y int) bool {
    return d.find(x) == d.find(y)
}

func (d *disjointSet) reset(x int) {
    d.parent[x] = x
    d.rank[x] = 0
}