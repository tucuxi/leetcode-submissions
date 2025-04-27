func minReorder(n int, connections [][]int) int {
    incoming := make([][]int, n)
    outgoing := make([][]int, n)
    for _, c := range connections {
        incoming[c[1]] = append(incoming[c[1]], c[0])
        outgoing[c[0]] = append(outgoing[c[0]], c[1])
    }
    reorient := 0
    visited := map[int]bool{0: true}
    for q := []int{0}; len(q) > 0; q = q[1:] {
        city := q[0]
        for _, next := range incoming[city] {
            if !visited[next] {
                q = append(q, next)
                visited[next] = true
            }
        }
        for _, next := range outgoing[city] {
            if !visited[next] {
                q = append(q, next)
                visited[next] = true
                reorient++
            }   
        }
    }
    return reorient
}