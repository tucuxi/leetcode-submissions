func distinctNames(ideas []string) int64 {
    var initials [26]map[string]bool
    for i := range initials {
        initials[i] = make(map[string]bool)
    }
    for _, s := range ideas {
        initials[s[0] - 'a'][s[1:]] = true
    }
    var count int64
    for i := range initials {
        for j := i + 1; j < len(initials); j++ {
            var mutuals int64
            for ideaA := range initials[i] {
                if initials[j][ideaA] {
                    mutuals++
                }
            }
            count += 2 * (int64(len(initials[i])) - mutuals) * (int64(len(initials[j])) - mutuals)
        }
    }
    return count
}