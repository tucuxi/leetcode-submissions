func areaOfMaxDiagonal(dimensions [][]int) int {
    a := 0
    d := 0
    for _, r := range dimensions {
        x, y := r[0], r[1]
        if h := x * x + y * y; h > d {
            d = h
            a = x * y
        } else if h == d && x * y > a {
            a = x * y
        }
    }
    return a
}