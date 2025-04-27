func minimumChairs(s string) int {
    chairs := 0
    maxChairs := 0
    for _, event := range s {
        if event == 'E' {
            chairs++
            maxChairs = max(maxChairs, chairs)
        } else {
            chairs--
        }
    }
    return maxChairs
}