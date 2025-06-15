class Solution {
    fun subtractProductAndSum(n: Int): Int {
        val s = n.toString()
        val prod = s.fold(1) { ac, ch -> ac * (ch - '0') }
        val sum = s.fold(0) { ac, ch -> ac + (ch - '0') }
        return prod - sum
    }
}