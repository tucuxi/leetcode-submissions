class Solution {
    fun complexNumberMultiply(num1: String, num2: String) =
        (Complex.of(num1) * Complex.of(num2)).toString()
}

data class Complex(val real: Int, val imag: Int) {
    
    companion object {
        fun of(num: String): Complex {
            val (real, imag) = num.subSequence(0, num.length - 1).split("+").map { it.toInt() }
            return Complex(real, imag)        
        }
    }
    
    operator fun times(other: Complex) =
        Complex(real * other.real - imag * other.imag, real * other.imag + imag * other.real)
        
    override fun toString() = "%d+%di".format(real, imag)
}