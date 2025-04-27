class Solution {
    fun numWaterBottles(numBottles: Int, numExchange: Int): Int {
        var drink = numBottles
        var empty = numBottles
        while (empty >= numExchange) {
            drink += empty / numExchange
            empty = empty / numExchange + empty % numExchange
        }
        return drink
    }
}