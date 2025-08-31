class Solution {
    fun getLeastFrequentDigit(n: Int): Int {
        return digits(n)
            .minWith(compareBy({ it.value }, { it.key }))
            .key
    }

    tailrec fun digits(n: Int, acc: MutableMap<Int, Int> = mutableMapOf()): Map<Int, Int> {
        return if (n == 0) {
            acc
        } else {
            val d = n % 10
            acc[d] = acc.getOrDefault(d, 0) + 1
            digits(n / 10, acc)
        } 
    }
}