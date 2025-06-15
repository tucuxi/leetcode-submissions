func removeSubfolders(folder []string) []string {
    var res []string

    slices.Sort(folder)
    res = append(res, folder[0])
    p := folder[0] + "/"
    for _, s := range folder[1:] {
        if !strings.HasPrefix(s, p) {
            res = append(res, s)
            p = s + "/"
        }
    }
    return res
}