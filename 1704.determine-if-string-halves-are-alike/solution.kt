class Solution {
    fun halvesAreAlike(s: String): Boolean {
        val split = s.length / 2
        return numberOfVowels(s.substring(0, split)) == numberOfVowels(s.substring(split))
    }
        
    fun numberOfVowels(s: String): Int =
        s.count { "aAeEiIoOuU".contains(it) }    
}