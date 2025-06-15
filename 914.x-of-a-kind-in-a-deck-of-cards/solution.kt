class Solution {
    fun hasGroupsSizeX(deck: IntArray): Boolean {
        
        fun gcd(a: Int, b: Int): Int =
            when {
                a == 0 -> b
                b == 0 -> a
                a > b -> gcd(b, a % b)
                else -> gcd(a, b % a)
            }

        return deck.groupBy { it }.map { it.value.size }.reduce { a, b -> gcd(a, b) } > 1
    }
}