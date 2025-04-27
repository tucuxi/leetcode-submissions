func memLeak(memory1 int, memory2 int) []int {
    i := 1
    for memory1 >= i || memory2 >= i {
        if memory2 > memory1 {
            memory2 -= i
        } else {
            memory1 -= i
        }
        i++
    }
    return []int{i, memory1, memory2}
}