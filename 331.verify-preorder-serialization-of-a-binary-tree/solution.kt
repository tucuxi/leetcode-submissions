class Solution {
    fun isValidSerialization(preorder: String): Boolean {
        var nodes = preorder.split(",")
        var index = 0
   
        fun validate(): Boolean {
            if (index == nodes.size) {
                return false
            }
            if (nodes[index++] == "#") {
                return true
            }
            return validate() && validate()
        }
        
        return validate() && index == nodes.size
    }
}