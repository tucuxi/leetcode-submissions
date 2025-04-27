func letterCombinations(digits string) []string {
    var (        
        res []string
        f func(int, []rune)
    )
    
    m := map[byte]string { 
        '2': "abc", 
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno", 
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    f = func(index int, prefix []rune) {
        if index == len(digits) {
            res = append(res, string(prefix))
            return
        }
        for _, c := range m[digits[index]] {
            f(index + 1, append(prefix, c))
        }
    }
    
    if len(digits) > 0 {
        f(0, nil)        
    }

    return res
}