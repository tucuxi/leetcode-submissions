func destCity(paths [][]string) string {
    cities := make(map[string]bool)
    for _, p := range paths {
        cities[p[1]] = true
    }
    for _, p := range paths {
        delete(cities, p[0])
    }
    for c := range cities {
        return c
    }
    return ""
}