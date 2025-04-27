type freq struct {
    char byte
    count []int
}

func rankTeams(votes []string) string {
    n := len(votes[0])

    teams := make(map[rune]int)
    for i, t := range votes[0] {
        teams[t] = i
    }

    positions := make([]*freq, n)
    for i := range positions {
        positions[i] = &freq{votes[0][i], make([]int, n)}
    }
    for _, v := range votes {
        for i := range v {
            j := strings.IndexByte(votes[0], v[i])
            positions[j].count[i]++
        }
    }

    slices.SortFunc(positions, func(a, b *freq) int {
        k := 0
        for k < n && a.count[k] == b.count[k] {
            k++
        } 
        if k == n {
            return int(a.char) - int(b.char)
        }
        return b.count[k] - a.count[k]
    })

    res := make([]byte, n)
    for i := range positions {
        res[i] = positions[i].char 
    }

    return string(res)
}