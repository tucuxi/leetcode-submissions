class Solution {
    fun multiply(num1: String, num2: String): String {
        if (num1 == "0" || num2 == "0") return "0"
        
        val res = IntArray(num1.length + num2.length)
        
        for (i in num1.lastIndex downTo 0) {
            for (j in num2.lastIndex downTo 0) {
                res[i + j + 1] += (num1[i] - '0') * (num2[j] - '0')
                res[i + j] += res[i + j + 1] / 10
                res[i + j + 1] %= 10
            }
        }
        
        val ans = StringBuilder()
        var i = 0
        while (res[i] == 0) i++
        while (i <= res.lastIndex) ans.append(res[i++].toString())
        
        return ans.toString()
    }
}