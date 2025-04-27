type entry struct {
    vertex int
    freq int
}

func secondMinimum(n int, edges [][]int, time int, change int) int {
    adj := make([][]int, n+1)
    for _, edge := range edges {
        i, j := edge[0], edge[1]
        adj[i] = append(adj[i], j)
        adj[j] = append(adj[j], i)
    }
    dist1 := make([]int, n+1)
    dist2 := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dist1[i] = -1
        dist2[i] = -1
    }
    dist1[1] = 0
    var q = []entry{ { 1, 1 } }
    for len(q) > 0 {
        vertex := q[0].vertex
        freq := q[0].freq
        q = q[1:]
        var timeTaken int
        if freq == 1 {
            timeTaken = dist1[vertex]
        } else {
            timeTaken = dist2[vertex]
        }
        if  timeTaken / change % 2 == 1 {
            timeTaken = change * (timeTaken / change + 1) + time
        } else {
            timeTaken += time
        }
        for _, neighbor := range adj[vertex] {
            if dist1[neighbor] == -1 {
                dist1[neighbor] = timeTaken
                q = append(q, entry{neighbor, 1})
            } else if dist2[neighbor] == -1 && dist1[neighbor] != timeTaken {
                if neighbor == n {
                    return timeTaken
                }
                dist2[neighbor] = timeTaken
                q = append(q, entry{neighbor, 2})
            }
        }
    }    
    return 0
}