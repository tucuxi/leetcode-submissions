func possibleStringCount(word string) int {
    res := 1
    var previous byte
    for i := range word {
        if word[i] != previous {
            previous = word[i]
        } else {
            res++
        }
    }
    return res
}