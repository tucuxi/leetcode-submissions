func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
    recipeSet := make(map[string]int)
    for r := range recipes {
        recipeSet[recipes[r]] = r
    }

    supplySet := make(map[string]bool)
    for _, supply := range supplies {
        supplySet[supply] = true
    }

    visited := make([]bool, len(recipes))
    canCreateRecipe := make(map[int]bool)
    var canCreate func(int) bool

    canCreate = func(r int) bool {
        if visited[r] {
            return canCreateRecipe[r]
        }

        visited[r] = true

        for _, ingredient := range ingredients[r] {
            fmt.Println(r, ingredient)
            if i, exists := recipeSet[ingredient]; exists {
                if !canCreate(i) {
                    canCreateRecipe[r] = false
                    return false
                }
            } else if !supplySet[ingredient] {
                canCreateRecipe[r] = false
                return false
            }
        }
        canCreateRecipe[r] = true
        return true
    }

    var res []string

    for r := range recipes {
        if canCreate(r) {
            res = append(res, recipes[r])
        }
    }

    return res 
}