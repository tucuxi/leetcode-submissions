func countStudents(students []int, sandwiches []int) int {
    var pref [2]int

    for _, st := range students {
        pref[st]++
    }
    for _, sa := range sandwiches {
        if pref[sa] == 0 {
            break
        }
        pref[sa]--
    }
    return pref[0] + pref[1]
}