func repairCars(ranks []int, cars int) int64 {
    return int64(sort.Search(math.MaxInt, func(m int) bool {
        sum := 0
        for _, r := range ranks {
            sum += int(math.Sqrt(float64(m) / float64(r)))
        }
        return sum >= cars
    }))
}