func wateringPlants(plants []int, capacity int) int {
    count := 0
    pos := -1
    rest := capacity
    for i := range plants {
        fmt.Println(pos, rest, count)
        for plants[i] > 0 {
            if rest >= plants[i] {
                count += i - pos
                pos = i
                rest -= plants[i]
                plants[i] = 0
            } else {
                count += pos + 1
                pos = -1
                rest = capacity
            }
        }
    }
    return count
}