func stoneGameIX(stones []int) bool {
    var r [3]int

    for _, stone := range stones {
        r[stone % 3]++
    }
    if r[0] % 2 == 0 {
        return r[1] >= 1 && r[2] >= 1
    }
    return r[1] - r[2] > 2 || r[2] - r[1] > 2
}