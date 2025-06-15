func nextGreatestLetter(letters []byte, target byte) byte {
    for _, l := range letters {
        if l > target {
            return l
        }
    }
    return letters[0]
}