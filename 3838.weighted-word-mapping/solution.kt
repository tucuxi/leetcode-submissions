class Solution {
    fun mapWordWeights(words: Array<String>, weights: IntArray): String {
        return words.joinToString("") { word ->
            val weight = word.sumOf { weights[it - 'a'] }
            val mappedChar = 'z' - weight % 26
                
           "$mappedChar"
        }
    }
}