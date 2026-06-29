class Solution {
    fun maxNumberOfBalloons(text: String): Int {
        val f = IntArray(26)
        text.forEach {
            f[it - 'a']++
        }
        return listOf(f['b'-'a'], f['a'-'a'], f['l'-'a'] / 2, f['o'-'a'] / 2, f['n'-'a']).min()!!
    }
}
