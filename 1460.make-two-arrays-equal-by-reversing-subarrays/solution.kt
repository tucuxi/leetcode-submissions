class Solution {
    fun canBeEqual(target: IntArray, arr: IntArray) =
        target.sorted() == arr.sorted()
}