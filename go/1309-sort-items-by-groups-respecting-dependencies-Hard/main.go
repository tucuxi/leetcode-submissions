func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
    m = augmentGroup(m, group)
    
    incomingGroup := make([]map[int]bool, m)
    outgoingGroup := make([]map[int]bool, m)
    for i := 0; i < m; i++ {
        incomingGroup[i] = make(map[int]bool)
        outgoingGroup[i] = make(map[int]bool)
    }
    
    incomingItem := make([]map[int]bool, n)
    outgoingItem := make([]map[int]bool, n)
    for i := 0; i < n; i++ {
        incomingItem[i] = make(map[int]bool)
        outgoingItem[i] = make(map[int]bool)
    }
    
    groupItems := make([][]int, m)
    
    for i := range beforeItems {
        g1 := group[i]
        groupItems[g1] = append(groupItems[g1], i)
        for _, b := range beforeItems[i] {
            g2 := group[b]
            if g2 == g1 {
                incomingItem[i][b] = true
                outgoingItem[b][i] = true
            } else {
                incomingGroup[g1][g2] = true
                outgoingGroup[g2][g1] = true                
            }
        }
    }
    
    res := make([]int, 0, n)

    eligibleGroups := filterEmpty(incomingGroup)
    for len(eligibleGroups) > 0 {
        g := eligibleGroups[0]
        eligibleGroups = eligibleGroups[1:]
        var eligibleItems []int
        for _, i := range groupItems[g] {
            if len(incomingItem[i]) == 0 {
                eligibleItems = append(eligibleItems, i)
            }
        }
        for len(eligibleItems) > 0 {
            i := eligibleItems[0]
            eligibleItems = eligibleItems[1:]
            res = append(res, i)
            for j := range outgoingItem[i] {
                delete(incomingItem[j], i)
                if len(incomingItem[j]) == 0 {
                    eligibleItems = append(eligibleItems, j)
                }
            }
        }
        for h := range outgoingGroup[g] {
            delete(incomingGroup[h], g)
            if len(incomingGroup[h]) == 0 {
                eligibleGroups = append(eligibleGroups, h)
            }
        }
    }

    if len(res) < n {
        return nil
    }
    return res
}

func augmentGroup(m int, group []int) int {
    for i := range group {
        if group[i] == -1 {
            group[i] = m
            m++
        }
    }
    return m
}

func filterEmpty(h []map[int]bool) []int {
    var res []int
    for i := range h {
        if len(h[i]) == 0 {
            res = append(res, i)
        }
    }
    return res
}