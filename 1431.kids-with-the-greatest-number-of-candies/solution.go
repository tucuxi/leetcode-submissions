func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := 0
    for _, c := range candies {
        if c > max {
            max = c
        }
    }
    res := make([]bool, 0, len(candies))
    for _, c := range candies {
        res = append(res, c + extraCandies >= max)
    }
    return res
}