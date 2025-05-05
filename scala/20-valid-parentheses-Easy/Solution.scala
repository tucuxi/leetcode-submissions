import scala.collection.mutable.Stack

object Solution {
    val validParen = Map('(' -> ')', '[' -> ']', '{' -> '}')
    def isValid(s: String): Boolean = {
        var paren = Stack[Char]()
        for (ch <- s) {
            if (validParen.contains(ch)) {
                paren.push(ch)
            } else if (paren.isEmpty) {
                return false
            } else {
                val pch = paren.pop
                if (ch != validParen(pch)) return false
            }
        }
        paren.isEmpty
    }
}