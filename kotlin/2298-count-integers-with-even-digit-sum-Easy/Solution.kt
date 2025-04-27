class Solution {
    fun countEven(num: Int): Int {
        var res = 0
        var n = num
        while (n > 0) {
            var sum = 0
            var m = n
            while (m > 0) {
                sum += m % 10
                m /= 10
            }
            if (sum % 2 == 0) {
                res++
            }
            n--
        }
        return res
    }
}