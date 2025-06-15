import kotlin.math.abs

class Solution {
    fun minMovesToSeat(seats: IntArray, students: IntArray): Int {
        return (seats.sorted() zip students.sorted()).fold(0) {
            acc, (seat, student) -> acc + abs(seat - student)
        }
    }
}