// this doesnt work for stupid reasons
func multiply(num1 string, num2 string) string {
    var atoi func(string) int64
    var itoa func(int64) string

    atoi = func(in string) int64 {
        var (
            sum int64 = 0
            place int64 = 1
        )
        for i := len(in) - 1; i >= 0; i-- {
            val := int64(in[i] - 48)
            sum += val * place
            place *= 10
        }
        return sum
    }

    int1 := atoi(num1)
    int2 := atoi(num2)

    product := int1 * int2

    itoa = func(in int64) string {
        if in == 0 {
            return "0"
        }

        var bytesReversed []byte
        var bytes []byte

        for in != 0 {
            rem := in % 10
            b := rem + 48
            bytesReversed = append(bytesReversed, byte(b))
            in /= 10
        }

        for i := len(bytesReversed) - 1; i >= 0; i-- {
            bytes = append(bytes, bytesReversed[i])
        }

        return string(bytes)
    }

    return itoa(product)
}

// this works nicely
func multiply(num1 string, num2 string) string {
    maxLength := len(num1) + len(num2)
    if maxLength == 0 {
        return "0"
    }

    digits := make([]int, maxLength)

    for i := 0; i < len(num1); i++ {
        for j := 0; j < len(num2); j++ {
            place := i + j

            digit1 := int(num1[len(num1) - 1 - i] - 48)
            digit2 := int(num2[len(num2) - 1 - j] - 48)

            carry := digit1 * digit2
            for carry != 0 && place < maxLength {
                sum := carry + digits[place]
                rem := sum % 10
                digits[place] = rem
                carry = sum / 10
                place++
            }
        }
    }

    var buffer bytes.Buffer
    nonZeroSeen := false
    for i := maxLength - 1; i >= 0; i-- {
        if !nonZeroSeen && digits[i] == 0 {
            if i == 0 {
                buffer.WriteByte(48)
            }
            continue
        }
        nonZeroSeen = true
        buffer.WriteByte(byte(digits[i] + 48))
    }

    return buffer.String()
}
