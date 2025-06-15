func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures))
    days := make([]int, 0, len(temperatures))
    for i, temperature := range temperatures {
        for j := len(days) - 1; j >= 0 && temperature > temperatures[days[j]]; j-- {
            day := days[j]
            res[day] = i - day
            days = days[:j]
        }
        days = append(days, i)
    }
    return res
}