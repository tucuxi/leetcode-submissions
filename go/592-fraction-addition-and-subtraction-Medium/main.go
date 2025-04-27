type fraction struct {
    nominator, denominator int
}

func add(a, b fraction) fraction {
    nominator := a.nominator * b.denominator + a.denominator * b.nominator
    denominator := a.denominator * b.denominator
    gcd := gcd(nominator, denominator)
    return fraction{nominator / gcd, denominator / gcd}
}

func (f fraction) String() string {
    return fmt.Sprintf("%d/%d", f.nominator, f.denominator)
}

func gcd(a, b int) int {
    if a < 0 {
        a = -a
    }
    if b < 0 {
        b = -b
    }
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

func fractionAddition(expression string) string {
    res := fraction{0, 1}
    p := 0

    scanNumber := func() (int, bool) {
        if p == len(expression) {
            return 0, false
        }
        number := 0
        sign := 1
        if expression[p] == '+' {
            p++
        } else if expression[p] == '-' {
            p++
            sign = -1
        }
        for ; p < len(expression) && expression[p] >= '0' && expression[p] <= '9'; p++ {
            number =  10 * number + int(expression[p] - '0')
        }
        return number * sign, true
    }

    scanFraction := func() (fraction, bool) {
        var ok bool
        var res fraction
        res.nominator, ok = scanNumber()
        if !ok {
            return res, false
        }
        p++
        res.denominator, ok = scanNumber()
        return res, true
    }

    for {
        f, ok := scanFraction()
        if !ok {
            break
        }
        res = add(res, f)
    }
    return res.String()
}