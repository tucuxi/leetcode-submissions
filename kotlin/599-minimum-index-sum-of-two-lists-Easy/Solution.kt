class Solution {
    fun findRestaurant(list1: Array<String>, list2: Array<String>): Array<String> {
        val common = list1
            .mapIndexedNotNull { i1, s1 ->
                list2.indexOf(s1).let { i2 ->
                    if (i2 == -1) {
                        null
                    } else {
                        Pair(i1, i2)
                    }
                } 
            }
            .sortedBy { (i1, i2) -> i1 + i2 }
        val minimumSum = with (common[0]) { first + second }
        return common
            .takeWhile { (i1, i2) -> i1 + i2 == minimumSum }
            .map { (i1, _) -> list1[i1] }
            .toTypedArray()
    }
}