func earliestFullBloom(plantTime []int, growTime []int) int {
    idx := make([]int, len(growTime))
    for i := range idx {
        idx[i] = i
    }
    sort.Slice(idx, func(i, j int) bool {
        return growTime[idx[i]] > growTime[idx[j]]
    })
    time, res := 0, 0
    for _, i := range idx {
        time += plantTime[i]
        if nt := time + growTime[i]; nt > res {
            res = nt
        }
    }
    return res
}