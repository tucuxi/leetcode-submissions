object Solution {
    def duplicateZeros(arr: Array[Int]): Unit = {
        val tmp = arr.clone
        var i = 0
        var j = 0
        while (j < arr.length) {
            arr(j) = tmp(i)
            if (tmp(i) == 0) {
                j += 1
                if (j < arr.length)
                    arr(j) = 0
            }
            j += 1
            i += 1
        }
    }
}