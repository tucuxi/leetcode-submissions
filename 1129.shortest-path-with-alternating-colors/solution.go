type edge struct{
    next int
    color int
}

type entry struct{
    node int
    dist int
    color int
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
    adj := make([][]edge, n)
    for _, redEdge := range redEdges {
        node := redEdge[0]
        adj[node] = append(adj[node], edge{redEdge[1], 0})
    }
    for _, blueEdge := range blueEdges {
        node := blueEdge[0]
        adj[node] = append(adj[node], edge{blueEdge[1], 1})
    }
    distColor := make([][2]int, n)
    for i := 1; i < n; i++ {
        distColor[i][0] = math.MaxInt
        distColor[i][1] = math.MaxInt
    }
    q := []entry{{0, 0, 0}, {0, 0, 1}}
    for len(q) > 0 {
        e := q[0]
        q = q[1:]
        for _, edge := range adj[e.node] {
            if edge.color != e.color && e.dist + 1 < distColor[edge.next][edge.color] {
                distColor[edge.next][edge.color] = e.dist + 1
                q = append(q, entry{edge.next, e.dist + 1, edge.color})
            }
        }
    }
    dist := make([]int, n)
    for i := 1; i < n; i++ {
        if distColor[i][0] == math.MaxInt && distColor[i][1] == math.MaxInt {
            dist[i] = -1
        } else if distColor[i][0] < distColor[i][1] {
            dist[i] = distColor[i][0]
        } else {
            dist[i] = distColor[i][1]
        }
    }
    return dist
}