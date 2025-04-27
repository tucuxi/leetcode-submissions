func numberOfWeakCharacters(properties [][]int) int {
    sort.Slice(properties, func(a, b int) bool {
        if properties[a][0] == properties[b][0] {
            return properties[a][1] < properties[b][1]
        }
        return properties[a][0] > properties[b][0]
    })
    res := 0
    max := -1
    for _, p := range properties {
        if p[1] < max {
            res++
        }
        if p[1] > max {
            max = p[1]
        }
    }
    return res
}