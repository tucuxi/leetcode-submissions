func equalSubstring(s string, t string, maxCost int) int {
    res := 0
    left := 0
    currentCost := 0
    for right := range(s) {
        currentCost += cost(int(s[right]), int(t[right]))
        for currentCost > maxCost {
            currentCost -= cost(int(s[left]), int(t[left]))
            left++
        }
        res = max(res, right - left + 1)
    }
    return res
}

func cost(a, b int) int {
    if a < b {
        return b-a
    }
    return a-b
}