func maximumCandies(candies []int, k int64) int {
    slices.Sort(candies)

    lo := 0
    hi := candies[len(candies) - 1]

    for lo < hi {
        candiesPerChild := lo + (hi - lo + 1) / 2
        children := int64(0)
        for i := len(candies) - 1; i >= 0; i-- {
            additionaLChildren := int64(candies[i] / candiesPerChild)
            if additionaLChildren == 0 {
                break
            }
            children += additionaLChildren
            if children >= k {
                break
            }
        }
        if children >= k {
            lo = candiesPerChild
        } else {
            hi = candiesPerChild - 1
        }
    }
    return lo
}