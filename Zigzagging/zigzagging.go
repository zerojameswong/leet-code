// my stupid version O(n^2)
func convert(s string, numRows int) string {
    straights := []byte{}
    diagonals := []byte{}

    converted := []byte{}

    if numRows == 1 {
        return s
    }

    diagonalLength := numRows - 2
    vLength := numRows + diagonalLength
    numVs := len(s) / vLength
    if len(s) % vLength > 0 {
        numVs += 1
    }

    for i := 0; i < len(s); i++ {
        if i % vLength < numRows {
            straights = append(straights, s[i])
        } else {
            diagonals = append(diagonals, s[i])
        }
    }
    
    for i := 0; i < numRows; i++ {
        for j := 0; j < numVs; j++ {
            if j * vLength + i > len(s) - 1 {
                continue
            }
            converted = append(converted, straights[j * numRows + i])
            if i != 0 && i != numRows - 1 {
                if (j + 1) * diagonalLength - i >= len(diagonals) {
                    continue
                }
                converted = append(converted, diagonals[(j + 1) * diagonalLength - i])
            }
        }
    }

    return string(converted)
}

// so much more elegant and O(n)
func convert(s string, numRows int) string {
    lines := make([]string, numRows)

    currentLine := 0
    direction := -1
    for i := range(s) {
        lines[currentLine] += s[i : i + 1]
        if direction == -1 {
            currentLine++
        } else {
            currentLine--
        }
        if currentLine == 0 || currentLine == numRows - 1 {
            direction *= -1
        }

    }

    result := ""
    for i := range(lines) {
        result += lines[i]
    }

    return result
}
