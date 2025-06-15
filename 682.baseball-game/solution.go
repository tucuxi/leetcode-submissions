func calPoints(ops []string) int {
    sc := []int{}
    for _, op := range ops {
        l := len(sc)-1
        switch op {
        case "+":
            sc = append(sc, sc[l] + sc[l-1])
        case "D":
            sc = append(sc, sc[l] * 2)
        case "C":
            sc = sc[:l]
        default:
            x, _ := strconv.Atoi(op)
            sc = append(sc, x)
        }
    }
    sum := 0
    for _, v := range sc {
        sum += v
    }
    return sum
}