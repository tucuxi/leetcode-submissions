func reorderLogFiles(logs []string) []string {
    j := len(logs) 
    for i := len(logs) - 1; i >= 0; i-- {
        k := strings.IndexByte(logs[i], ' ') + 1
        if ch := logs[i][k]; '0' <= ch && ch <= '9' {
            j--
            logs[i], logs[j] = logs[j], logs[i]
        } 
    }
    sort.Slice(logs[:j], func(i, j int) bool {
        li := logs[i]
        lj := logs[j]
        ki := strings.IndexByte(li, ' ') + 1
        kj := strings.IndexByte(lj, ' ') + 1
        return li[ki:] < lj[kj:] || li[ki:] == lj[kj:] && li < lj
    })
    return logs
}