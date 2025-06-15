func hammingDistance(x int, y int) int {
    count := 0
    for z := x ^ y; z != 0; z &= z - 1 {
        count++
    }
    return count
}