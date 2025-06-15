func checkIfPangram(sentence string) bool {
    occ := [26]bool{}
    rem := len(occ)
    for _, c := range sentence {
        if !occ[c - 'a'] {
            occ[c - 'a'] = true
            rem--
        }
    }
    return rem == 0
}