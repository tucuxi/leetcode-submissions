func losingPlayer(x int, y int) string {
    t := min(x, y / 4)
    return [2]string{"Bob", "Alice"}[t & 1]
}