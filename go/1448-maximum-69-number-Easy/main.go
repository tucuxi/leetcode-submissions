func maximum69Number(num int) int {
    s := strconv.Itoa(num)
    t := strings.Replace(s, "6", "9", 1)
    res, _ := strconv.Atoi(t)
    return res
}