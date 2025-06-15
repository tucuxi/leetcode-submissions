class Solution {
    fun squareIsWhite(coordinates: String): Boolean {
        return ((coordinates[0] - 'a') + (coordinates[1] - '1')) % 2 != 0
    }
}