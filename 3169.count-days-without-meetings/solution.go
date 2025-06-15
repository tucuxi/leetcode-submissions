type Change struct {
    day int
    value int
}

func countDays(days int, meetings [][]int) int {
    var changes []Change
    
    for _, m := range meetings {
        changes = append(changes, Change{m[0], 1}, Change{m[1], -1})
    }
    slices.SortFunc(changes, func(a, b Change) int { return a.day - b.day })

    count := 0
    day := 1
    res := 0

    for i := 0; i < len(changes); i++ {
        currentDay := changes[i].day
        currentChange := changes[i].value
        for i + 1 < len(changes) && changes[i + 1].day == currentDay {
            currentChange += changes[i + 1].value
            i++
        }
        if count == 0 {
            res += currentDay - day
        }
        count += currentChange
        day = currentDay + 1
    }

    res += days - day + 1
    return res
}