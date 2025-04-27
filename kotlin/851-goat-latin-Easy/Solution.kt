class Solution {
    fun toGoatLatin(sentence: String): String =
        sentence
            .split(' ')
            .mapIndexed { i, w ->
                val s = if (w.first() in "aAeEiIoOuU") {
                    w
                } else {
                    w.drop(1) + w.first() 
                }
                s + "ma" + "a".repeat(i + 1)
            }
            .joinToString(" ")
}