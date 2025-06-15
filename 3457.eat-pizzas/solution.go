func maxWeight(pizzas []int) int64 {
    slices.Sort(pizzas)

    res := int64(0)
    k := len(pizzas)
    evenDays := k/8
    oddDays := k/4 - evenDays

    for range oddDays {
        k--
        res += int64(pizzas[k])
    }
    for range evenDays {
        k -= 2
        res += int64(pizzas[k])
    }

    return res
}