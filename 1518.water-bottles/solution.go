func numWaterBottles(numBottles int, numExchange int) int {
    return rec(numBottles, numBottles, numExchange)
}

func rec(drunken, emptyBottles, numExchange int) int {
    drink := emptyBottles / numExchange
    if drink == 0 {
        return drunken
    }
    rest := emptyBottles % numExchange
    return rec(drink + drunken, drink + rest, numExchange)
}