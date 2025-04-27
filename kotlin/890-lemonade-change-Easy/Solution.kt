class Solution {
    fun lemonadeChange(bills: IntArray): Boolean {
        var num5 = 0
        var num10 = 0
        for (bill in bills) {
            when (bill) {
                5 -> num5++
                10 -> {
                    if (num5 == 0) return false
                    num5--
                    num10++
                }
                20 -> {
                    if (num10 == 0) {
                        if (num5 < 3) return false
                        num5 -= 3
                    } else {
                        if (num5 == 0) return false
                        num5--
                        num10--
                    }
                }
            }
        }
        return true
    }
}