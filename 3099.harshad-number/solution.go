func sumOfTheDigitsOfHarshadNumber(x int) int {
    s := 0
    for y := x; y > 0; y /= 10 {
        s += y % 10
    }
    if x % s > 0 {
        return -1
    }
    return s
}