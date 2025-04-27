type Position struct {
    x,y int
    keys int
}

func shortestPathAllKeys(grid []string) int {
    keysCount := 0
    var lastQ []Position

    isKey := func(x, y int) bool {
        return grid[x][y] >= 'a' && grid[x][y] <= 'k'
    }
    
    m := len(grid)
    n := len(grid[0])
    for i, row := range grid {
        for j, cell := range row {
            if cell == '@' {
                lastQ = []Position{Position{x:i, y:j, keys:0}}
            } else if isKey(i,j) {
                keysCount++
            }
        }
    }
    
    seen := make(map[Position]bool)
    q := make([]Position, 0)
    steps := 0
    for len(lastQ) > 0 {
        for _, pos := range lastQ {
            x,y := pos.x, pos.y
            //check if all keys found
            if pos.keys == ((1<<keysCount)-1){
                return steps-1
            }
            if x < 0 || x == m || y < 0 || y == n { // out of bounds
                continue
            } else if grid[x][y] == '#' { // Wall
                continue
            } else if seen[pos] { // Already seen
                continue
            }
    
            //lock
            if grid[x][y] >= 'A' && grid[x][y] <= 'K'{
                if pos.keys & (1 << int(grid[x][y] -'A')) == 0 {
                    //don't have the key
                    continue
                }
            }
            //key
            if isKey(x, y){
                pos.keys |= 1 << int(grid[x][y] - 'a')
            }

            seen[pos] = true
            q = append(q, Position{x:x+1, y:y, keys:pos.keys})
            q = append(q, Position{x:x-1, y:y, keys:pos.keys})
            q = append(q, Position{x:x, y:y+1, keys:pos.keys})
            q = append(q, Position{x:x, y:y-1, keys:pos.keys})
        }
        lastQ, q = q, lastQ[:0]
        steps++
    }
    return -1
}