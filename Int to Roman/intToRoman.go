func intToRoman(num int) string {
    runes := [][]string{
        []string{"I", "V", "X"},
        []string{"X", "L", "C"},
        []string{"C", "D", "M"},
        []string{"M", "", ""},
    }

    digit := 0
    roman := ""
    
    for num != 0 {
        lastDigit := num % 10
        num = num / 10
        digit += 1

        runeSet := runes[digit - 1]

        prepend := ""
        if lastDigit <= 3 {
            prepend += strings.Repeat(runeSet[0], lastDigit)
        } else if lastDigit <= 5 {
            prepend += strings.Repeat(runeSet[0], 5 - lastDigit)
            prepend += runeSet[1]
        } else if lastDigit <= 8 {
            prepend += runeSet[1]
            prepend += strings.Repeat(runeSet[0], lastDigit - 5)
        } else {
            prepend += runeSet[0]
            prepend += runeSet[2]
        }
        roman = prepend + roman
    }

    return roman
}