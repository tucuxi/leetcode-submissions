func areAlmostEqual(s1 string, s2 string) bool {
    diff := []int{}
    for i := range s1 {
        if s1[i] != s2[i] {
            diff = append(diff, i)
        }
    }
    return len(diff) == 0 ||
        len(diff) == 2 && s1[diff[0]] == s2[diff[1]] && s1[diff[1]] == s2[diff[0]]
}