func checkDistances(s string, distance []int) bool {
    occ := [26]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
    for i := range s {
        k := s[i] - 'a'
        if occ[k] < 0 {
            occ[k] = i
        } else if i - occ[k] != distance[k] + 1 {
            return false
        }
    }
    return true
}