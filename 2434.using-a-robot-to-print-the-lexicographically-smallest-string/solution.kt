class Solution {
    fun robotWithString(s: String): String {
        return buildString {
            val count = s.groupingBy { it }.eachCount().toMutableMap()
            val stack = ArrayDeque<Char>()
            var minCharacter = 'a'

            s.forEach { character ->
                stack.addLast(character)
                count[character] = count.getValue(character) - 1
                while (minCharacter < 'z' && count.getOrDefault(minCharacter, 0) == 0) {
                    minCharacter++
                }
                while (stack.isNotEmpty() && stack.last() <= minCharacter) {
                    append(stack.removeLast())
                }
            }
        }
    }
}