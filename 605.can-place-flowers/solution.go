func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {
        return true
    }
    empty := 1
    for _, plot := range flowerbed {
        if plot == 0 {
            empty++
            if empty == 3 {
                n--
                if n == 0 {
                    return true
                }
                empty = 1
            }
        } else {
            empty = 0
        }
    }
    return n == 1 && empty == 2
}