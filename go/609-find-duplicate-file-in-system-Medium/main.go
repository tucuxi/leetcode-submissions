type entry struct {
    path string
    next *entry
}

func findDuplicate(paths []string) [][]string {
    d := map[string]*entry{}
    for _, p := range paths {
        i := 0
        for ; p[i] != ' '; i++ {
        }
        pre := p[:i]
        i++
        for i < len(p) {
            b := i
            for ; p[i] != '('; i++ {
            }
            file := p[b:i]
            i++
            c := i
            for ; p[i] != ')'; i++ {
            }
            content := p[c:i]
            nextEntry := d[content]
            d[content] = &entry{pre + "/" + file, nextEntry}
            i += 2
        }
    }
    res := [][]string{}
    for _, v := range d {
        if v.next == nil {
            continue
        }
        dupes := []string{}
        for e := v; e != nil; e = e.next {
            dupes = append(dupes, e.path)
        }
        res = append(res, dupes)
    }
    return res
}