func findTheCity(n int, edges [][]int, distanceThreshold int) int {
    dist := make([][]int, n)
    for i := range dist {
        dist[i] = make([]int, n)
        for j := range n {
            dist[i][j] = 1e6
        }
        dist[i][i] = 0
    }
    for _, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        dist[u][v] = w
        dist[v][u] = w
    }
    for k := range n {
        for i := range n {
            for j := range n {
                dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
            }
        }
    }
    minCities := math.MaxInt
    resCity := -1
    for i := range n {
        cities := 0
        for j := range n {
            if dist[i][j] > 0 && dist[i][j] <= distanceThreshold {
                cities++
            }
        } 
        if cities <= minCities {
            minCities = cities
            resCity = i
        }
    }
    return resCity
}