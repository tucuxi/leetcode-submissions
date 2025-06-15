func subdomainVisits(cpdomains []string) []string {
    domainVisits := make(map[string]int)
    for _, s := range cpdomains {
        var count int
        var domain string
        fmt.Sscanf(s, "%d %s", &count, &domain)
        for {
            domainVisits[domain] += count
            i := 0
            for i < len(domain) && domain[i] != '.' {
                i++
            }
            if i == len(domain) {
                break
            }
            domain = domain[i + 1:]
        }
    }
    res := make([]string, 0, len(domainVisits))
    for domain, count := range domainVisits {
        res = append(res, fmt.Sprintf("%d %s", count, domain))
    }
    return res
}