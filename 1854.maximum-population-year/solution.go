func maximumPopulation(logs [][]int) int {
    const base = 1950
    pop := make([]int, 101)
    for _, p := range(logs) {
        pop[p[0] - base]++
        pop[p[1] - base]--
    }
    off := 0
    for y := 1; y < len(pop); y++ {
        pop[y] += pop[y - 1]
        if (pop[y] > pop[off]) {
            off = y
        }
    }
    return base + off
}