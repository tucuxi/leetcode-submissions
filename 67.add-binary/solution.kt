import java.math.BigInteger

class Solution {
    fun addBinary(a: String, b: String): String {
        val a1 = binaryToBigInteger(a)
        val b1 = binaryToBigInteger(b)
        val c = a1.add(b1)
        return bigIntegerToBinary(c)
    }
    
    fun binaryToBigInteger(a: String): BigInteger =
        a.fold(BigInteger.ZERO) { res, digit ->
            val doubled = res.shiftLeft(1)
            if (digit == '0') doubled else doubled.inc() 
        }
        
    fun bigIntegerToBinary(n: BigInteger): String {
        if (n == BigInteger.ZERO) {
            return "0"
        }
        val res = StringBuilder()
        var m = n
        while (m != BigInteger.ZERO) {
            if (m.mod(BigInteger("2")) == BigInteger.ZERO) {
                res.append('0')
            } else {
                res.append('1')
            }
            m = m.shiftRight(1)
        }
        return res.reverse().toString()
    }
}