func commonChars(words []string) []string {
    var containedInAll [26]int

    for i := range containedInAll {
        containedInAll[i] = math.MaxInt
    }

    for _, w := range words {
        var containedInCurrent [26]int

        for _, r := range w {
            containedInCurrent[r - 'a']++
        }
        for i := range containedInCurrent {
            containedInAll[i] = min(containedInAll[i], containedInCurrent[i])
        }
    }

    var res []string

    for i := range containedInAll {
        for range containedInAll[i] {
            res = append(res, fmt.Sprintf("%c", i + 'a'))
        }
    }

    return res
}