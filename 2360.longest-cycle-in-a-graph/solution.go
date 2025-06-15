func longestCycle(edges []int) int {
    lc := -1
    k := 0
    visited := make([]int, len(edges))
    for i := range edges {
        k1 := k
        for j := edges[i]; j != -1; j = edges[j] {
            k++
            if visited[j] > 0 {
                if visited[j] > k1 && k - visited[j] > lc {
                    lc = k - visited[j]
                }
                break
            }
            visited[j] = k
        }
    }    
    return lc
}