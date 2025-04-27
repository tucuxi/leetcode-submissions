func timeRequiredToBuy(tickets []int, k int) int {
    t := 0
    for i := 0; ; i = (i+1) % len(tickets) {
        if tickets[i] > 0 {
            tickets[i]--
            t++
            if i == k && tickets[i] == 0 {
                return t
            }
        }
    }
}