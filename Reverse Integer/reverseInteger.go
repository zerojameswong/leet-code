// my horribly unwieldy way to do it
func reverse(x int) int {
    stringified := strconv.Itoa(x)
    reversed := []byte{}

    sign := 1
    if (x >> 31) & 1 == 1 {
        sign = -1
    }

    if sign == 1 {
        max := strconv.Itoa(math.MaxInt32)
        if len(max) == len(stringified) {
            for i := 0; i < len(stringified); i++ {
                if stringified[len(stringified) - 1 - i] > max[i] {
                    return 0
                }
            }
        }
    }

    if sign == -1 {
        min := strconv.Itoa(math.MinInt32)
        if len(min) == len(stringified) {
            for i := 1; i < len(stringified); i++ {
                if stringified[len(stringified) - 1 - i] > min[i] {
                    return 0
                }
            }
        }

        reversed = append(reversed, stringified[0])
        stringified = stringified[1:]
    }

    for i := len(stringified) - 1; i >= 0; i-- {
        if stringified[i : i + 1] != "0" {
            reversed = append(reversed, stringified[i])
        }
    }

    reversedStringified, _ := strconv.Atoi(string(reversed))

    return reversedStringified
}

// insane concept from youtube
// we cant store a 32 bit overflow but we can store an integer that is one decimal number shorter than the max / min integer
// then we can use this in tandem with the last digit to determine if the reversed integer would overflow 32 bits
// i.e. either the first n-1 digits would cause an overflow by themselves or if they dont but the nth will
func reverse(x int) int {
    reversed := 0

    maxComparison := math.MaxInt32 / 10
    maxLast := math.MaxInt32 % 10
    minComparison := math.MinInt32 / 10
    minLast := math.MinInt32 % 10

    for x != 0 {
        remainder := x % 10
        x = x / 10

        if reversed > maxComparison || reversed == maxComparison && remainder > maxLast {
            return 0
        }

        if reversed < minComparison || reversed == minComparison && remainder < minLast {
            return 0
        }

        reversed = reversed * 10 + remainder
    }

    return reversed
}
