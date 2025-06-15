func carPooling(trips [][]int, capacity int) bool {
    h := make(map[int]int)
    for _, trip := range trips {
        h[trip[1]] += trip[0]
        h[trip[2]] -= trip[0]
    }
    
    changes := make([][2]int, 0, len(h))
    for location, change := range h {
        changes = append(changes, [2]int{location, change})
    }
    
    sort.Slice(changes, func(i, j int) bool { return changes[i][0] < changes[j][0] })
    
    s := 0
    for _, change := range changes {
        s += change[1]
        if s > capacity {
            return false
        }
    }
    return true
}