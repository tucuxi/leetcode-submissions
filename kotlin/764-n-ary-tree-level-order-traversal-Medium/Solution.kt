/**
 * Definition for a Node.
 * class Node(var `val`: Int) {
 *     var children: List<Node?> = listOf()
 * }
 */

class Solution {
    fun levelOrder(root: Node?): List<List<Int>> {
        val levels = mutableListOf<MutableList<Int>>()
        
        fun recurse(node: Node?, level: Int) {
            node?.apply {
                if (level > levels.lastIndex) {
                    levels.add(mutableListOf<Int>())
                } 
                levels[level].add(`val`)
                children.forEach { recurse(it, level  + 1) }
            }
        }
        
        recurse(root, 0)
        return levels
    }
}