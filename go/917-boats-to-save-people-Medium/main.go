func numRescueBoats(people []int, limit int) int {
    res := 0
    sort.Ints(people)
    for p, q := 0, len(people) - 1; q >= p; q-- {
        if people[p] + people[q] <= limit {
            p++
        }
        res++
    }
    return res
}