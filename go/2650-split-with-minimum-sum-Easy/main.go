func splitNum(num int) int {
    b := []byte(strconv.Itoa(num))
    if len(b) % 2 == 1 {
        b = append(b, '0')
    }
    sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
    num1, num2 := 0, 0
    for i := 0; i < len(b); i += 2 {
        num1 = (num1 * 10) + (int(b[i] - '0'))
        num2 = (num2 * 10) + (int(b[i + 1] - '0'))
    }
    return num1 + num2
}