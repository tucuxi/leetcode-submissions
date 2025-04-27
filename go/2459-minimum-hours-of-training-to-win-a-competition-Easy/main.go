func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
    training := 0
    totalEnergy := 0
    for _, eng := range energy {
        totalEnergy += eng
    }
    if initialEnergy <= totalEnergy {
        training += totalEnergy - initialEnergy + 1
    }
    currentExperience := initialExperience
    for _, exp := range experience {
        if currentExperience <= exp {
            missing := exp - currentExperience + 1
            currentExperience += missing
            training += missing
        }
        currentExperience += exp
    } 
    return training
}