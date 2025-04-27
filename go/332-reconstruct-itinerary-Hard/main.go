func findItinerary(tickets [][]string) []string {
    graph := make(map[string][]string)
    for _, t := range tickets {
        graph[t[0]] = append(graph[t[0]], t[1])
    }
    for _, c := range graph {
        sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })
    }    

    var names []string
    
    node := "JFK"
    path := []string{node}
    for len(path) > 0 {
        if len(graph[node]) > 0 {
            path = append(path, node)
            node, graph[node] = graph[node][0], graph[node][1:]
        } else {
            names = append(names, node)
            node = path[len(path) - 1]
            path = path[:len(path) - 1]
        }
    }
    for i, j := 0, len(names) - 1; i < j; i, j = i + 1, j - 1 {
        names[i], names[j] = names[j], names[i]
    }
    return names
}