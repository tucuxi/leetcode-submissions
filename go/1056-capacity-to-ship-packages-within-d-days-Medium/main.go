func shipWithinDays(weights []int, days int) int {
    w := 0
    for _, weight := range weights {
        w += weight
    }
    return sort.Search(w, func(i int) bool { return valid(weights, days, i) })
}

func valid(weights []int, days int, capacity int) bool {
    w := 0
    d := 0
    for _, weight := range weights {
        if weight > capacity {
            return false
        }
        if w + weight > capacity {
            d++
            if d == days {
                return false
            }
            w = weight
        } else {
            w += weight
        }
    }   
    return true
}