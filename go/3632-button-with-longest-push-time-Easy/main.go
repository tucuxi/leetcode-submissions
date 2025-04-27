func buttonWithLongestTime(events [][]int) int {
    longestButtonTime := 0
    longestButtonIndex := -1
    previousTime := 0
    for _, event := range events {
        index, time := event[0], event[1]
        thisButtonTime := time - previousTime
        if thisButtonTime > longestButtonTime {
            longestButtonTime = thisButtonTime
            longestButtonIndex = index
        } else if thisButtonTime == longestButtonTime && index < longestButtonIndex {
            longestButtonIndex = index
        }
        previousTime = time
    }
    return longestButtonIndex
}