func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
    var filteredRestaurants [][]int

    for _, r := range restaurants {
        if (r[2] == 1 || veganFriendly == 0) && r[3] <= maxPrice && r[4] <= maxDistance {
            filteredRestaurants = append(filteredRestaurants, r)
        }
    }

    slices.SortFunc(filteredRestaurants, func(a, b []int) int {
        if a[1] == b[1] {
            return b[0] - a[0]
        }
        return b[1] - a[1]
    })

    res := make([]int, 0, len(filteredRestaurants))

    for _, f := range filteredRestaurants {
        res = append(res, f[0])
    }

    return res
}