class Solution {
    fun pyramidTransition(bottom: String, allowed: List<String>): Boolean {
        val allowedMap: Map<String, String> = allowed
            .groupingBy { pattern -> pattern.substring(0, 2) }
            .aggregate { _, acc, pattern, _ -> acc.orEmpty() + pattern.substring(2) }
        
        fun dfs(bottom: String, above: String = ""): Boolean {
            if (bottom.length == 1 && above.length == 0) {
                return true
            }
            if (bottom.length <= 1) {
                return dfs(above)
            }
            val currentBlock = bottom.substring(0, 2)
            allowedMap[currentBlock]?.forEach { replacement ->
                if (dfs(bottom.substring(1), above + replacement)) {
                    return true
                }
            }
            return false
        }

        return dfs(bottom)
    }
}