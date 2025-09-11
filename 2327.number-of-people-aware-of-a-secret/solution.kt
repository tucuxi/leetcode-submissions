import java.math.BigInteger

class Solution {
    fun peopleAwareOfSecret(n: Int, delay: Int, forget: Int): Int {
        val know = Array<BigInteger>(n) { BigInteger.ZERO }.apply { set(0, BigInteger.ONE) }
        var shared = BigInteger.ZERO
        var total = BigInteger.ONE

        for (day in delay until forget) {
            shared += know[day - delay]
            total += shared
            know[day] = shared
        }
        for (day in forget until n) {
            shared += know[day - delay] - know[day -forget]
            total += shared - know[day - forget]
            know[day] = shared
        }

        return total.mod(BigInteger("1000000007")).toInt()
    }
}