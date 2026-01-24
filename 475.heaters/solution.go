func findRadius(houses []int, heaters []int) int {
    res := 0

    slices.Sort(heaters)

    for _, house := range houses {
        i, _ := slices.BinarySearch(heaters, house)

        distance := 0

        if i == 0 {
            distance = heaters[0] - house
        } else if i == len(heaters) {
            distance = house - heaters[len(heaters) - 1]
        } else {
            distance = max(distance, min(house - heaters[i - 1], heaters[i] - house))
        }

        res = max(res, distance)
    }

    return res
}