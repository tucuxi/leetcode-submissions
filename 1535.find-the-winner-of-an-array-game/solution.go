func getWinner(arr []int, k int) int {
    max := arr[0]
    for _, a := range arr {
        if a > max {
            max = a
        }
    }
    cur := arr[0]
    run := 0
    for i := 1; i < len(arr); i++ {
        if cur > arr[i] {
            run++
        } else {
            cur = arr[i]
            run = 1
        }
        if run == k || cur == max {
            return cur
        }
    }
    panic("no solution")
}