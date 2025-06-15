class Solution {
    fun areNumbersAscending(s: String): Boolean {
        var p = 0
        for (t in s.split(' ')) {
            if (t.first().isDigit()) {
                val num = t.toInt()
                if (num <= p) {
                    return false
                }
                p = num
            }
        }
        return true
    }
}