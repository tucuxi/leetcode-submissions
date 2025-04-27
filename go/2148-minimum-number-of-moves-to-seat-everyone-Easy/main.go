func minMovesToSeat(seats []int, students []int) int {
    slices.Sort(seats)
    slices.Sort(students)
    res := 0
    for i := range seats {
        res += abs(seats[i] - students[i])
    }
    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}