func largestInteger(num int) int {
    s := strconv.Itoa(num)
    a := [2][]byte{}
    for i := range s {
        p := s[i] % 2
        a[p] = append(a[p], s[i])
    }
    sort.Slice(a[0], func(i, j int) bool { return a[0][i] > a[0][j] })
    sort.Slice(a[1], func(i, j int) bool { return a[1][i] > a[1][j] })
    var sb strings.Builder
    k := [2]int{}
    for i := range s {
        p := s[i] % 2
        sb.WriteByte(a[p][k[p]])
        k[p]++
    }
    res, _ := strconv.Atoi(sb.String())
    return res
}