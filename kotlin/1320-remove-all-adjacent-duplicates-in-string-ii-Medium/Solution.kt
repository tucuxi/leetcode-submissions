class Solution {
    fun removeDuplicates(s: String, k: Int): String {
        val sb = StringBuilder()
        var sk = 1
        var cp = '$'
        s.forEachIndexed { i, ci ->
            sb.append(ci)
            if (ci == cp) {
                sk++
            } else {
                sk = 1
            }
            cp = ci
            if (sk == k) {
                with(sb) {
                    setLength(length - k)
                    sk = 0
                    if (isNotEmpty()) {
                        var j = lastIndex
                        cp = get(j)
                        while (j >= 0 && get(j) == cp) {
                            sk++
                            j--
                        }    
                    }
                }
            } 
        }
        return sb.toString()
    }
}