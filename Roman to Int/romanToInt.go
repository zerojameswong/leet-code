func romanToInt(s string) int {
    charToVal := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    sum := 0

    for i := range(s) {
        if i + 1 < len(s) && charToVal[s[i]] < charToVal[s[i + 1]] {
            sum -= charToVal[s[i]]
        } else {
            sum += charToVal[s[i]]
        }
    }

    return sum
}
