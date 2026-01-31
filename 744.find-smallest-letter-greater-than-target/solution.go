func nextGreatestLetter(letters []byte, target byte) byte {
    i, _ := slices.BinarySearch(letters, target + 1)
    return letters[i % len(letters)]
}