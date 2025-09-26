class Solution {
    fun fractionToDecimal(numerator: Int, denominator: Int): String {
        val res = StringBuilder()
        
        if (numerator < 0 && denominator > 0 || numerator > 0 && denominator < 0) {
            res.append('-')
        }

        val n = abs(numerator.toLong())
        val d = abs(denominator.toLong())

        res.append(n / d)

        var r = n % d * 10

        if (r == 0L) {
            return res.toString()
        }

        res.append('.')

        val positionByRest = mutableMapOf<Long, Int>()

        while (r != 0L && r !in positionByRest) {
            positionByRest[r] = res.length
            res.append(r / d)
            r = r % d * 10
        }

        if (r != 0L) {
            res.insert(positionByRest.getValue(r), '(')
            res.append(')')
        }

        return res.toString()
    }
}