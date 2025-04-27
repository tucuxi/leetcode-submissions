func corpFlightBookings(bookings [][]int, n int) []int {
    flights := make([]int, n + 1)
    for _, b := range bookings {
        first, last, seats := b[0] - 1, b[1], b[2]
        flights[first] += seats
        flights[last] -= seats
    }
    for i := 1; i <= n; i++ {
        flights[i] += flights[i - 1]
    }
    return flights[:n]
}