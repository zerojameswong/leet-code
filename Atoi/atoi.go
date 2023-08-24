func myAtoi(s string) int {
    sign := 1
    doneSigning := false
    startedReading := false
    result := 0

    digits := []int{}

    for i := range(s) {
        // if started reading and non number encountered
        if startedReading && (s[i] < 48 || s[i] > 57) {
            break
        }
        // whitespace
        if s[i] == 32 {
            continue
        }
        // positive
        if s[i] == 43 && !doneSigning {
            doneSigning = true
            continue
        }
        // negative
        if s[i] == 45 && !doneSigning {
            doneSigning = true
            sign = -1
            continue
        }
        // digits
        if s[i] >= 48 && s[i] <= 57  {
            startedReading = true
            digits = append(digits, sign * (int(s[i]) - 48))
        }
    }

    maxIntComparison := math.MaxInt32 / 10
    maxLast := math.MaxInt32 % 10
    minIntComparison := math.MinInt32 / 10
    minLast := math.MinInt32 % 10

    for i := range(digits) {
        if result > maxIntComparison || result == maxIntComparison && digits[i] > maxLast {
            result = math.MaxInt32
            break
        }

        if result < minIntComparison || result == minIntComparison && digits[i] < minLast {
            result = math.MinInt32
            break
        }
        result = result * 10 + digits[i]
    }

    return result
}