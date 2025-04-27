func removeAlmostEqualCharacters(word string) int {
    var (
        res int
        prev rune
    )
    
    for _, r := range word {
        if r == prev || r+1 == prev || r-1 == prev {
            res++
            prev = 0
        } else {
            prev = r
        }
    }
    
    return res
}