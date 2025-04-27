/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
    total := 0
    
    emap := map[int]*Employee{}
    for _, e := range employees {
        emap[e.Id] = e
    }
 
    var dfs func(int)
    dfs = func(id int) {
        e := emap[id]
        total += e.Importance
        for _, sub := range e.Subordinates {
            dfs(sub)
        }
    }
    
    dfs(id)
    return total
}