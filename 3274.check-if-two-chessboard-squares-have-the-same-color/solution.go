func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
    return (coordinate1[0] - 'a' + coordinate1[1] - '1') & 1 == (coordinate2[0] - 'a' + coordinate2[1] - '1') & 1
}