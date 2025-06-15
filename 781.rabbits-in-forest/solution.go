func numRabbits(answers []int) int {
    h := make(map[int]int)
    for _, a := range answers {
        h[a]++
    }

    res := 0
    for a, c := range h {
        res += (a + c) / (a + 1) * (a + 1) 
    }

    return res
}