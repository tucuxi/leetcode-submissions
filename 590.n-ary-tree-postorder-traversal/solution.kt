/**
 * Definition for a Node.
 * class Node(var `val`: Int) {
 *     var children: List<Node?> = listOf()
 * }
 */

class Solution {
    fun postorder(root: Node?): List<Int> =
        root?.run { children.flatMap { postorder(it) } + `val` } ?: emptyList()
}