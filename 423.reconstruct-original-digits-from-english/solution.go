func originalDigits(s string) string {
    f := make(map[rune]int)
    for _, ch := range s {
        switch ch {
        case 'f': f['5']++
        case 'g': f['8']++
        case 'i': f['9']++
        case 'o': f['1']++
        case 'r': f['3']++
        case 's': f['7']++
        case 'u': f['4']++
        case 'w': f['2']++
        case 'x': f['6']++
        case 'z': f['0']++
        }
    }
    f['1'] -= f['0'] + f['2'] + f['4']
    f['3'] -= f['0'] + f['4']
    f['5'] -= f['4']
    f['7'] -= f['6']
    f['9'] -= f['5'] + f['6'] + f['8']
    
    var res strings.Builder
    for ch := '0'; ch <= '9'; ch++ {
        for i := 0; i < f[ch]; i++ {
            res.WriteRune(ch)
        }
    }
    return res.String()
}