func numOfUnplacedFruits(fruits []int, baskets []int) int {
    res := 0

    for _, f := range fruits {
        j := 0
        for j < len(baskets) && baskets[j] < f {
            j++
        }
        if j < len(baskets) {
            baskets = append(baskets[:j], baskets[j + 1:]...)
        } else {
            res++
        }
    }

    return res
}