type occurrence struct {
    character rune
    number int
}

func repeatLimitedString(s string, repeatLimit int) string {
    h := make(map[rune]int)
    for _, r := range s {
        h[r]++
    }

    var occurrences []occurrence
    
    for character, number := range h {
        occurrences = append(occurrences, occurrence{character, number})
    }

    slices.SortFunc(occurrences, func(a, b occurrence) int {
        return int(b.character - a.character)
    })

    var previousCharacter rune
    var sb strings.Builder

    outer:
    for len(occurrences) > 0 {
        if occurrences[0].number == 0 {
            occurrences = occurrences[1:]
            continue
        }
        nextIndex := 0
        n := min(repeatLimit, occurrences[nextIndex].number)
        if occurrences[nextIndex].character == previousCharacter {
            n = 1
            for {
                nextIndex++
                if nextIndex == len(occurrences) {
                    break outer
                }
                if occurrences[nextIndex].number >= n {
                    break
                }
            }
        }
        for range n {
            sb.WriteRune(occurrences[nextIndex].character)
        }
        occurrences[nextIndex].number -= n
        previousCharacter = occurrences[nextIndex].character
    }
    return sb.String()
}