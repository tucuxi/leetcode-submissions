class Solution {
    fun timeRequiredToBuy(tickets: IntArray, k: Int): Int {
        var t = 0
        var i = 0
        while (true) {
            if (tickets[i] > 0) {
                tickets[i]--
                t++
            }
            if (tickets[i] == 0 && i == k) {
                return t
            }
            if (++i == tickets.size) {
                i = 0
            }
        }
    }
}