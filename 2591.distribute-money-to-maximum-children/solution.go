func distMoney(money int, children int) int {
    m := money - children
    if m < 0 {
        return -1
    }
    c := m / 7
    switch {
        case c == children && money % 8 == 0:
            return c
        case c >= children:
            return children - 1
        case c == children - 1 && money % 8 == 4:
            return c - 1
        default:
            return c
    }
}