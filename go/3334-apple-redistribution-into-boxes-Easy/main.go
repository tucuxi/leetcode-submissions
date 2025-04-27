func minimumBoxes(apple []int, capacity []int) int {
    totalApples := 0
    for _, a := range apple {
        totalApples += a
    }
    slices.Sort(capacity)
    i := len(capacity)
    for totalApples > 0 {
        i--
        totalApples -= capacity[i]
    }
    return len(capacity) - i
}