class Solution {
    fun totalMoney(n: Int): Int {
        var res = 0
        repeat(n) { i -> res += i/7 + i%7 + 1}
        return res
    }
}