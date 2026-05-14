class Solution {
    fun sortVowels(s: String): String {
        val f = s
            .filter { "aeiou".indexOf(it) >= 0 }
            .groupingBy { it }
            .eachCount()

        val v = f.entries
            .sortedWith(
                compareBy(
                    { -it.value },
                    { s.indexOf(it.key) }
                )
            )
            .joinToString("") { (ch, n) -> "$ch".repeat(n) }

        var j = 0
        
        val res = StringBuilder()
        
        for (i in s.indices) {
            if ("aeiou".indexOf(s[i]) >= 0) {
                res.append(v[j])
                j++
            } else {
                res.append(s[i])
            }
        }
        
        return res.toString()
    }
}