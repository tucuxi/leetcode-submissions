func numSteps(s string) int {
    b := make([]byte, 0, len(s) + 1)
    for i := len(s) - 1; i >= 0; i-- {
        b = append(b, s[i])
    }
    k := 0
    for len(b) > 1 {
        k++
        c := b[0] & 1
        if c == 0 {
            b = b[1:]
        } else {
            for i := range b {
                sum := b[i] + c
                b[i] = sum & ^byte(2)
                c = (sum & 2) >> 1
                if c == 0 {
                    break
                }
            }
            if c > 0 {
                b = append(b, '1')
            }
        }
      
    }
    return k
}