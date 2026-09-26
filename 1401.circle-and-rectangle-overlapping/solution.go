func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
    px, py := xCenter, yCenter

    if xCenter < x1 {
        px = x1
    } else if xCenter > x2 {
        px = x2
    }
    if yCenter < y1 {
        py = y1
    } else if yCenter > y2 {
        py = y2
    }

    dx := xCenter - px
    dy := yCenter - py

    return dx * dx + dy * dy <= radius * radius
}