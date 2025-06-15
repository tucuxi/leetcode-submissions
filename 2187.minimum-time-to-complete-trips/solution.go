func minimumTime(time []int, totalTrips int) int64 {
    return int64(sort.Search(time[0] * totalTrips, func(i int) bool {
        trips := 0
        for _, t := range time {
            trips += i / t
        }
        return trips >= totalTrips
    }))
}