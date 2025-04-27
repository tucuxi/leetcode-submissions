func numFriendRequests(ages []int) int {
    var h, p [121]int

    for _, age := range ages {
        h[age]++
    }
    
    for age := 1; age < len(p); age++ {
        p[age] = p[age - 1] + h[age]
    }

    res := 0

    for _, age := range ages {
        if minAge := (age + 14) / 2; minAge < age {
            res += p[age] - p[minAge] - 1
        }
    }

    return res
}