func maxLength(arr []string) int {
    return recurse(arr, "")
}

func recurse(arr []string, pre string) int {
    if len(arr) == 0 {
        return len(pre)
    }
    len := recurse(arr[1:], pre)
    newPre := pre + arr[0]
    if uniqueCharacters(newPre) {
        lenWith := recurse(arr[1:], newPre)
        if lenWith > len {
            return lenWith
        }
    }
    return len
}

func uniqueCharacters(s string) bool {
    f := [26]bool{}
    for i := range s {
        k := s[i] - 'a'
        if f[k] {
            return false
        }
        f[k] = true
    }
    return true
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}