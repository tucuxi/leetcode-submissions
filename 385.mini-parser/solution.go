/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
    ni, _ := parse(s)
    return ni
}

func parse(s string) (*NestedInteger, string) {
    if s[0] == '-' || s[0] >= '0' && s[0] <= '9' {
        return parseInteger(s)
    }
    if s[0] == '[' {
        return parseList(s)
    }
    panic(fmt.Sprintf("invalid string: %s", s))
}

func parseInteger(s string) (*NestedInteger, string) {
    sign, value := 1, 0
    if s[0] == '-' {
        sign = -sign
        s = s[1:]
    }
    for len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
        value = 10 * value + int(s[0] - '0')
        s = s[1:]
    }
    ni := &NestedInteger{}
    ni.SetInteger(sign * value)
    return ni, s
}

func parseList(s string) (*NestedInteger, string) {
    fmt.Printf("parseList(%s)\n", s)
    list := &NestedInteger{}
    s = eat(s, '[')
    if s[0] != ']' {
        var ni *NestedInteger
        ni, s = parse(s)
        list.Add(*ni)
        for len(s) > 0 && s[0] == ',' {
            s = eat(s, ',')
            ni, s = parse(s)
            list.Add(*ni)
        }
    }
    s = eat(s, ']')
    return list, s
}

func eat(s string, b byte) string {
    if len(s) == 0 {
        panic(fmt.Sprintf("expected %c, got empty string", b))
    }
    if s[0] != b {
        panic(fmt.Sprintf("expected %c, got string %s", b, s))
    }
    return s[1:]
}