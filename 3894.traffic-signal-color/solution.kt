class Solution {
    fun trafficSignal(timer: Int): String =
        when {
            timer == 0 ->
                "Green"
            timer == 30 ->
                "Orange"
            30 < timer && timer <= 90 ->
                "Red"
            else ->
                "Invalid"
        }
}