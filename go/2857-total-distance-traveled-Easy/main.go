func distanceTraveled(mainTank int, additionalTank int) int {
    km := 0
    for mainTank >= 5 {
        km += 50
        mainTank -= 5
        if additionalTank > 0 {
            additionalTank--
            mainTank++
        }
    }
    km += 10 * mainTank
    return km
}