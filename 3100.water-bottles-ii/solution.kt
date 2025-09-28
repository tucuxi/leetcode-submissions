class Solution {
    fun maxBottlesDrunk(numBottles: Int, numExchange: Int): Int {
        var res = numBottles
        var empty = numBottles
        var exchange = numExchange
        while (empty >= exchange) {
            empty = empty - exchange + 1
            res++
            exchange++
        }
        return res
    }
}