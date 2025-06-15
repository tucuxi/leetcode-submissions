func reverse(x int) int {
    r := []byte(strconv.Itoa(x))
    i := 0
    if r[0] == '-' {
        i++
    }
    for j := len(r) - 1; j > i; i, j = i + 1, j -1 {
        r[i], r[j] = r[j], r[i]
    }
    res, err := strconv.ParseInt(string(r), 10, 32)
    if err == nil {
        return int(res)
    }
    return 0
}