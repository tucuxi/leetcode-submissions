class Solution {
    fun areaOfMaxDiagonal(dimensions: Array<IntArray>): Int {
        return dimensions
            .maxWith(
                compareBy(
                    { (l, w) -> l * l + w * w },
                    { (l, w) -> l * w }
                )
            )
            .let { (l, w) -> l * w }
    }
}