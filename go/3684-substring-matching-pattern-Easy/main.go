func hasMatch(s string, p string) bool {
    star := strings.IndexByte(p, '*')
    i := strings.Index(s, p[:star])
    j := strings.LastIndex(s, p[star + 1:])
    return i != -1 && j != -1 && (i + star <= j || star == 0 || star == len(p) - 1)
}