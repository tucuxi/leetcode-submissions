func scoreValidator(events []string) []int {
    score := 0
    counter := 0

    for _, e := range events {
        switch e {
            case "W": counter++
            case "WD", "NB": score++
            default: score += int(e[0]-'0')
        }
        if counter == 10 {
            break
        }
    }
    return []int{score, counter}
}