func maxSatisfaction(satisfaction []int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(satisfaction)))
    sum, ltc := 0, 0
    for _, s := range satisfaction {
        sum += s
        if sum < 0 {
            break
        }
        ltc += sum
    }
    return ltc
}