func countVowelSubstrings(word string) int {
    res := 0
    for i := range word {
        fa, fe, fi, fu, fo := false, false, false, false, false
        inner: for j := i; j < len(word); j++ {
            switch word[j] {
                case 'a': fa = true
                case 'e': fe = true
                case 'i': fi = true
                case 'o': fo = true
                case 'u': fu = true
                default: break inner
            }
            if fa && fe && fi && fo && fu {
                res++
            }
        }
    }
    return res
}