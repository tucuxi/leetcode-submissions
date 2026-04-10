func decodeCiphertext(encodedText string, rows int) string {
    n := len(encodedText)
    cols := n / rows
    res := make([]byte, 0, n)

    for startCol := 0; startCol < cols; startCol++ {
        row, col := 0, startCol

        for row < rows && col < cols {
            index := row*cols + col
            res = append(res, encodedText[index])
            row++
            col++
        }
    }

    end := len(res)
    for end > 0 && res[end-1] == ' ' {
        end--
    }

    return string(res[:end])
}