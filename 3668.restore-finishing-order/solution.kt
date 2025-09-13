class Solution {
    fun recoverOrder(order: IntArray, friends: IntArray): IntArray {
        val friendsSet = friends.toSet()
        return order.filter { friendsSet.contains(it) }.toIntArray()
    }
}