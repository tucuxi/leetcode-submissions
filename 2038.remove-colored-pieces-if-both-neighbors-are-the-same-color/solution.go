func winnerOfGame(colors string) bool {
    aCluster := make(map[int]int)
    bCluster := make(map[int]int)
    for i := 0; i < len(colors); {
        j := advance(colors, 'A', i)
        if j - i >= 3 {
            aCluster[j - i]++
        }
        k := advance(colors, 'B', j)
        if k - j >= 3 {
            bCluster[k - j]++
        }
        i = k
    }
    return turns(aCluster) > turns(bCluster)
}

func advance(colors string, color byte, i int) int {
    for i < len(colors) && colors[i] == color {
        i++
    }
    return i
}

func turns(cluster map[int]int) int {
    res := 0
    for length, count := range cluster {
        if length >= 3 {
            res += count * (length - 2)
        }
    }
    return res
}