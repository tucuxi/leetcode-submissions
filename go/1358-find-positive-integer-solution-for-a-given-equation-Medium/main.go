/** 
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) [][]int {
    var res [][]int
    
    outer:
    for x := 1; x <= 1000; x++ {
        for y := 1; y <= 1000; y++ {
            f := customFunction(x, y)
            if f == z {
                res = append(res, []int{x, y}) 
            }
            if f >= z {
                continue outer
            }
        }
    }
    
    return res
}