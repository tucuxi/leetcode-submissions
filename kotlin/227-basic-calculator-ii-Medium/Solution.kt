class Solution {
    fun calculate(s: String): Int {
        val scanner = Scanner(s)
  
        fun factor(): Long {
            val token = scanner.take() as Num
            return token.value.toLong()
        }
        
        fun term(): Long {
            var a = factor()
            var token = scanner.peek()
            while (token is Multiply || token is Divide) {
                scanner.take()
                val b = factor()
                a = if (token is Multiply) a * b else a / b
                token = scanner.peek()
            }
            return a
        }

        fun expr(): Long {
            var a = term()
            var token = scanner.peek()
            while (token is Add || token is Subtract) {
                scanner.take()
                val b = term()
                a = if (token is Add) a + b else a - b
                token = scanner.peek()
            }
            return a
        }
        
        val res = expr()
        return res.toInt()
    }
    
}

class Scanner(val text: String) {
    private var offset = 0
    private var next = scan()
    
    fun peek(): Token = next
    
    fun take(): Token = next.also { next = scan() }
    
    private fun scan(): Token {
        while (offset < text.length && text[offset] == ' ') {
            offset++
        }
        if (offset == text.length) {
            return End()
        }
        var ch = text[offset++]
        return when (ch) {
            '+' -> Add()
            '-' -> Subtract()
            '*' -> Multiply()
            '/' -> Divide()
            in '0'..'9' -> {
                var value = ch - '0'
                while (offset < text.length && text[offset].isDigit()) {
                    value = 10 * value + (text[offset] - '0')
                    offset++
                }
                Num(value)
            }
            else -> Error()
        }
    }    
}

open class Token
open class Op : Token()
class Add : Op()
class Subtract : Op()
class Multiply: Op()
class Divide: Op()
class Num(val value: Int) : Token()
class End : Token()
class Error : Token()