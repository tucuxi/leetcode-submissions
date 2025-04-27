func countValidWords(sentence string) int {
    res := 0
    for _, w := range strings.Fields(sentence) {
        var digits, hyphens, punctuations int
        for _, c := range w {
            switch c {
                case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
                    digits++
                case '-':
                    hyphens++
                case ',', '.', '!':
                    punctuations++
            }
        }
        if digits > 0 || hyphens > 1 || punctuations > 1 {
            continue
        }
        if hyphens == 1 {
            p := strings.Index(w, "-")
            if p == 0 || p == len(w) - 1 || w[p - 1] < 'a' || w[p - 1] > 'z' || w[p + 1] < 'a' || w[p + 1] > 'z' {
                continue
            }
        }
        if punctuations == 1 {
            p := strings.IndexAny(w, ",.!")
            if p != len(w) - 1 {
                continue
            }
        }
        res++
    }
    return res
}