func numberOfAlternatingGroups(colors []int) int {
    res := 0
    
    for i := range colors {
        j := (i + 1) % len(colors)
        k := (i + 2) % len(colors)
        if colors[i] != colors[j] && colors[j] != colors[k] {
            res++
        }
    }

    return res
}