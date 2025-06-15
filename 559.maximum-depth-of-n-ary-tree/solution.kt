/**
 * Definition for a Node.
 * class Node(var `val`: Int) {
 *     var children: List<Node?> = listOf()
 * }
 */

class Solution {
    fun maxDepth(root: Node?): Int =
        if (root == null) 0 else (root.children.map { maxDepth(it) }.max() ?: 0) + 1
}