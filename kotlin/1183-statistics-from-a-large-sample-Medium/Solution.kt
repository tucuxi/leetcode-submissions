class Solution {
    fun sampleStats(count: IntArray): DoubleArray {
        val total = count.sum()
        
        fun minimum(): Double {
            var i = 0
            while (count[i] == 0) {
                i++
            }
            return i.toDouble()
        }
        
        fun maximum(): Double {
            var i = count.lastIndex
            while (count[i] == 0) {
                i--
            }
            return i.toDouble()
        }
        
        fun mean(): Double {
            var sum = 0.0
            for (i in 0..count.lastIndex) {
                sum += i.toDouble() * count[i].toDouble()
            }
            return sum / total  
        }
        
        fun median(): Double {
            var acc = 0
            var a = -1
            while (true) {
                val acc2 = acc + 2 * count[a + 1]
                if (acc2 > total) break
                acc = acc2
                a++
            }
            // Postcondition: count.slice(0..a).sum() <= total
            if (a == -1) {
                return 0.0
            }
            var b = a
            do {
                if (++b == count.size) {
                    return a.toDouble() 
                }  
            } while (count[b] == 0)
            if (acc == total) {
                while (count[a] == 0) {
                    a--
                }
                return (a + b).toDouble() / 2
            }
            return b.toDouble()
        }
        
        fun mode(): Double {
            var maxCount = count[0]
            var maxIndex = 0
            for (i in 1..count.lastIndex) {
                if (count[i] > maxCount) {
                    maxCount = count[i]
                    maxIndex = i
                }
            }
            return maxIndex.toDouble()
        }
        
        return doubleArrayOf(minimum(), maximum(), mean(), median(), mode())
    }
}