const sep = "-"

func convertDateToBinary(date string) string {
    var res []string
    for _, s := range strings.Split(date, sep) {
        n, _ := strconv.ParseUint(s, 10, 16)
        res = append(res, strconv.FormatUint(n, 2))
    }
    return strings.Join(res, sep)
}