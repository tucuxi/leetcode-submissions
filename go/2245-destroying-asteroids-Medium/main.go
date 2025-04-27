func asteroidsDestroyed(mass int, asteroids []int) bool {
    destroyed := true
    for i := 0; destroyed && i < len(asteroids); {
        destroyed = false
        for j := i; j < len(asteroids); j++ {
            if asteroids[j] <= mass {
                mass += asteroids[j]
                asteroids[i], asteroids[j] = asteroids[j], asteroids[i]
                i++
                destroyed = true
            }
        }
    }
    return destroyed
}