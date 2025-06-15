func canReach(s string, minJump int, maxJump int) bool {
    lastIndex := len(s) - 1 
    r := make([]bool, len(s))
    r[0] = true
    j := 0
    for i := range r {
        if r[i] {
            for j = max(j, i + minJump); j <= min(i + maxJump, lastIndex); j++ {
                if s[j] == '0' {
                    r[j] = true
                }
            }
        }
    }
    return r[lastIndex]
}