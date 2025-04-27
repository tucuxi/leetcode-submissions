class Solution {
    fun countConsistentStrings(allowed: String, words: Array<String>) =
        words.count { it.isConsistent(allowed) }    
}

fun String.isConsistent(allowed: String) = all { it in allowed }