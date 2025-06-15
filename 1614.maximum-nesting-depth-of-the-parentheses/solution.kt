class Solution {
    fun maxDepth(s: String): Int {
        var max = 0
        var current = 0
        s.forEach { 
            if (it == '(') 
                max = Math.max(++current, max)
            else if (it == ')')
                --current
        }
        return max
    }
}