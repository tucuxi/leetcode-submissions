class Solution {
    fun numberOfBeams(bank: Array<String>): Int {
        var res = 0
        var previous = 0
        bank.forEach { row ->
            val devices = row.count { it == '1' }
            if (devices > 0) {
                res += previous * devices
                previous = devices
            }
        }
        return res
    }
}