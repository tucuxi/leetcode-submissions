class Solution {
    fun reorderedPowerOf2(n: Int): Boolean {
        val digitsOfN = internal(n)
        return (0..31).any { b ->
            val candidate = internal(1 shl b)
            digitsOfN.contentEquals(candidate)
        }
    }

    fun internal(n: Int): IntArray {
        return IntArray(10).apply {
            var k = n
            while (k > 0) {
                this[k % 10]++
                k /= 10
            }
        }
    }
}