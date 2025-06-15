const (
    bulkyDimension = 10_000
    bulkyVolume = 1_000_000_000
    heavyMass = 100
    
)
func categorizeBox(length int, width int, height int, mass int) string {
    bulky := length >= bulkyDimension || width >= bulkyDimension || height >= bulkyDimension ||
        int64(length) * int64(width) * int64(height) >= bulkyVolume
    heavy := mass >= heavyMass
    switch {
        case bulky && heavy: return "Both"
        case bulky: return "Bulky"
        case heavy: return "Heavy"
        default: return "Neither"
    }
}