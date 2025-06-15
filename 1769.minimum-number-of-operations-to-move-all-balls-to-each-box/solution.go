func minOperations(boxes string) []int {
    res := make([]int, len(boxes))
    preSum := 0
    preCnt := 0
    posSum := 0
    posCnt := 0
    for i := len(boxes) - 1; i >= 0; i-- {
        posSum += posCnt
        if boxes[i] == '1' {
            posCnt++
        }
    }
    for i := range boxes {
        preSum += preCnt
        res[i] = preSum + posSum
        if boxes[i] == '1' {
            preCnt++
            posCnt--
        }
        posSum -= posCnt
    } 
    return res
}

/*

         1  1  0
         -------
posCnt   2  1  0
posSum   1  0  0
preCnt   1  2  2
preSum   0  1  3

        1  1  0  0  1  0  1
        -------------------
posCnt  4  3  2  2  2  1  1
posSum  11 8  6  4  2  1  0

*/