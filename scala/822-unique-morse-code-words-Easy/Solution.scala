object Solution {
    val morse = Array(".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..",
                      ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.",
                      "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..")
    
    def uniqueMorseRepresentations(words: Array[String]): Int = {
        words.map(w => w.map(c => morse(c - 'a')).mkString).distinct.size    
    }
}