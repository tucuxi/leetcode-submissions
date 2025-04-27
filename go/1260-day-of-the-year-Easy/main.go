func dayOfYear(date string) int {
    v := strings.Split(date, "-")
    y, _ := strconv.Atoi(v[0])
    m, _ := strconv.Atoi(v[1])
    d, _ := strconv.Atoi(v[2])
    res := d
    for i := 1; i < m; i++ {
        switch i {
        case 1, 3, 5, 7, 8, 10, 12:
            res += 31
        case 4, 6, 9, 11:
            res += 30
        case 2:
            if y % 4 == 0 && y != 1900 {
                res += 29
            } else {
                res += 28
            }
        }
    }
    return res
}