func numBusesToDestination(routes [][]int, source int, target int) int {
    if source == target {
        return 0
    }
    h := make(map[int][]int)
    for bus, route := range routes {
        for _, stop := range route {
            h[stop] = append(h[stop], bus)
        } 
    }
    res := 0
    v := make([]bool, len(routes))
    q := make([]int, 0, len(h[source]))
    for _, bus := range h[source] {
        v[bus] = true
        q = append(q, bus)
    }
    for len(q) > 0 {
        res++
        for i := len(q); i > 0; i-- {
            bus := q[0]
            q = q[1:]
            v[bus] = true
            for _, stop := range routes[bus] {
                if stop == target {
                    return res
                }
                for _, nextBus := range h[stop] {
                    if !v[nextBus] {
                        v[nextBus] = true
                        q = append(q, nextBus)
                    }
                }
            }
        }
    }
    return -1
}