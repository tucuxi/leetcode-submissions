func numberOfSubstrings(s string) int {
    res := 0
    pos := []int{-1, -1, -1}    
    i := 0

    for ; i < len(s); i++ {
        pos[s[i] - 'a'] = i
        if pos[0] != -1 && pos[1] != -1 && pos[2] != -1 {
            break
        }
    }
    for ; i < len(s); i++ {
        pos[s[i] - 'a'] = i
        res += 1 + min(pos[0], pos[1], pos[2])        
    }

    return res
}