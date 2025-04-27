class Solution {
    fun destCity(paths: List<List<String>>): String {
        val startCities = paths.map { it[0] }
        val destinationCities = paths.map { it[1] }
        val res = destinationCities - startCities
        return res.first()
    }
}