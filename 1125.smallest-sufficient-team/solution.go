type skillSet uint16
type skillTranslation map[string]skillSet
type peopleSet uint64

func smallestSufficientTeam(reqSkills []string, people [][]string) []int {
    translator := makeSkillTranslation(reqSkills)
    peopleSkills := translator.translateAll(people)
    targetSkills := translator.translate(reqSkills)
    
    dp := make([]peopleSet, 1 << len(reqSkills))
    for i := 1; i < len(dp); i++ {
        dp[i] = (1 << len(people)) - 1
    }
    
    var targetSkillsMask skillSet
    
    for {
        targetSkillsMask++
        for i, personSkillsMask := range peopleSkills {
            smallerSkillsMask := targetSkillsMask & ^personSkillsMask
            if smallerSkillsMask != targetSkillsMask {
                peopleMask := dp[smallerSkillsMask] | (1 << i)
                if peopleMask.size() < dp[targetSkillsMask].size() {
                    dp[targetSkillsMask] = peopleMask
                }
            }
        }
        if targetSkillsMask == targetSkills {
            break
        }
    }
    
    res := make([]int, 0)
    for i := range people {
        if dp[targetSkills] & (1 << i) != 0 {
            res = append(res, i)
        }
    } 
    return res
}

func makeSkillTranslation(ss []string) skillTranslation {
    res := make(skillTranslation)
    for i, s := range ss {
        res[s] = 1 << i
    }
    return res
}

func (st skillTranslation) translate(ss []string) skillSet {
    var res skillSet
    for _, s := range ss {
        res |= st[s]
    }
    return res
}

func (st skillTranslation) translateAll(all [][]string) []skillSet { 
    res := make([]skillSet, len(all))
    for i, ss := range all {
        res[i] = st.translate(ss)
    }
    return res
}

func (ps peopleSet) size() int {
    res := 0
    for ps != 0 {
        ps &= ps - 1
        res++
    }
    return res
}