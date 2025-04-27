const (
    leftAndRight = 0b011_1111_1100
    center =       0b000_1111_0000
    onlyLeft =     0b000_0011_1100
    onlyRight =    0b011_1100_0000
)

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
    reservedRows := make(map[int]int)
    for _, s := range reservedSeats {
        reservedRows[s[0]] |= 1 << s[1]
    }
    groups := 2 * (n - len(reservedRows))
    for _, row := range reservedRows {
        switch {
            case row & leftAndRight == 0:
                groups += 2
            case row & onlyLeft == 0 || row & center == 0 || row & onlyRight == 0:
                groups++
        }
    }
    return groups
}