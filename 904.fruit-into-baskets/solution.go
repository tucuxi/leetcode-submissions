func totalFruit(fruits []int) int {
    basket := map[int]int{}
    max, cur := 0, 0
    for l, r := 0, 0; r < len(fruits); r++ {
        basket[fruits[r]]++
        cur++
        if len(basket) <= 2 {
            if cur > max {
                max = cur
            }
        } else {
            for {
                t := fruits[l]
                basket[t]--
                cur--
                l++
                if basket[t] == 0 {
                    delete(basket, t)
                    break
                }
            }
        }
    }
    return max
}