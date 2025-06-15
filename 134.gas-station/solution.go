func canCompleteCircuit(gas []int, cost []int) int {
    balance := 0
    tank := 0
    start := 0
    for i := range gas {
        balance += gas[i] - cost[i]
        tank += gas[i] - cost[i]
        if tank < 0 {
            tank = 0
            start = i + 1
        }
    }
    if balance >= 0 {
        return start
    }
    return -1
}
