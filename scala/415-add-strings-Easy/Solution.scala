object Solution {
    def addStrings(num1: String, num2: String): String = {

        def digitOrElse(num: String, index: Int, defaultValue: Int = 0): Int =
            if (index < 0 || index >= num.length) defaultValue else num(index).asDigit

        def toChar(n: Int): Char = (n + '0').toChar
        
        val rev1 = num1.reverse
        val rev2 = num2.reverse
        var acc = 0
        val res = new StringBuilder()
        for (i <- 0 until math.max(rev1.length, rev2.length)) {
            val sum = digitOrElse(rev1, i) + digitOrElse(rev2, i) + acc
            res += toChar(sum % 10)
            acc = sum / 10
        }
        if (acc > 0) res += toChar(acc) 
        res.result.reverse
    }
}