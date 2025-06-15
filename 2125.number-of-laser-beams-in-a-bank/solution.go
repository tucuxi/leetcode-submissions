func numberOfBeams(bank []string) int {
    res := 0
    prev := 0
    for _, row := range bank {
        devices := 0
        for _, cell := range row {
            if cell == '1' {
                devices++
            }
        }
        if devices > 0 {
            res += prev * devices
            prev = devices
        }
    }
    return res
}