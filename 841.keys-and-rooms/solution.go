func canVisitAllRooms(rooms [][]int) bool {
    const initial = 0
    unlocked := map[int]bool{initial: true}
    for queue := []int{initial}; len(queue) > 0; queue = queue[1:] {
        current := queue[0]
        for _, key := range rooms[current] {
            if !unlocked[key] {
                queue = append(queue, key)
                unlocked[key] = true
            }
        }
    }
    return len(unlocked) == len(rooms)
}