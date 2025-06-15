func twoEditWords(queries []string, dictionary []string) []string {
    var res []string

    queries:
    for _, q := range queries {

        words:
        for _, w := range dictionary {
            if len(q) != len(w) {
                continue
            }
            d := 0
            for i := range q {
                if q[i] != w[i] {
                    d++
                    if d > 2 {
                        continue words
                    }
                }
            }
            res = append(res, q)
            continue queries
        }
    }

    return res
}