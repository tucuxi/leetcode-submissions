class Solution {
    fun smallestNumber(n: Int, t: Int): Int {
        var i = n

        while (digitsProduct(i) % t != 0) {
            i++
        }
        return i  
    }

    fun digitsProduct(n: Int): Int {
        var p = 1
        var r = n

        while (r > 0) {
            p *= r % 10
            r /= 10
        }
        return p
    }
}