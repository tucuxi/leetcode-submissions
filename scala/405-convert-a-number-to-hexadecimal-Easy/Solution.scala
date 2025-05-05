object Solution {
    def toHex(num: Int): String = {
        if (num == 0) return "0"
        val sb = new StringBuilder(8)
        var n = num
        while (n != 0) {
            val digit = toHexDigit(n & 15)
            sb.append(digit)
            n >>>= 4
        }
        sb.reverse.result
    }
    
    def toHexDigit(digit: Int): Char = {
        if (digit < 10) ('0' + digit).toChar else ('a' - 10 + digit).toChar
    }
}