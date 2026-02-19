class Solution {
    fun addBinary(a: String, b: String): String {
        var i = a.lastIndex
        var j = b.lastIndex
        var c = 0

        val reversed = buildString {
            while (i >= 0 || j >= 0 || c > 0) {
                if (i >= 0) {
                    c += a[i--] - '0'
                }
                if (j >= 0) {
                    c += b[j--] - '0'
                }
                append('0' + (c and 1))
                c = c shr 1
            }
        }

        return reversed.reversed().toString()
    }
}