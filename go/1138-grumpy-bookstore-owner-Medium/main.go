func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    satisfied := 0
    runningExtra := 0
    maxExtra := 0
    for i := range grumpy {
        satisfied += (1 - grumpy[i]) * customers[i]
        if i >= minutes {
            runningExtra -= grumpy[i - minutes] * customers[i - minutes]
        }
        runningExtra += grumpy[i] * customers[i] 
        if runningExtra > maxExtra {
            maxExtra = runningExtra
        }
    }
    return satisfied + maxExtra
}