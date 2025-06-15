func dividePlayers(skill []int) int64 {
    var chemistry int64
    var count [1001]int
    minSkill := math.MaxInt
    maxSkill := math.MinInt
    for _, s := range skill {
        count[s]++
        minSkill = min(minSkill, s)
        maxSkill = max(maxSkill, s)
    }
    target := minSkill + maxSkill
    for s := minSkill; s <= target / 2; s++ {
        if s == target - s {
            if count[s] % 2 == 1 {
                return -1
            }
            chemistry += int64(count[s] / 2 * s * s)
        } else {
            if count[s] != count[target - s] {
                return -1
            }
            chemistry += int64(count[s] * s * (target - s))
        }
    }
    return chemistry
}
