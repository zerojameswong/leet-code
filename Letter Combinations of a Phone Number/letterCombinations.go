// this one got heated and i took it to reddit and got the solution which was that you have to deep copy rather than just append or else things get funky
// O(n4^n)
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    
    lettersPerDigit := map[byte]string {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    var (
        old, new [][]byte
    )
    old = [][]byte{[]byte("")}

    for i := range(digits) {
        letters := lettersPerDigit[digits[i]]
        for j := range(letters) {
            for k := range(old) {
                newEntry := make([]byte, len(old[k]))
                copy(newEntry, old[k])
                newEntry = append(newEntry, letters[j])
                new = append(new, newEntry)
            }
        }
        old = new
        new = [][]byte{}
    }

    ret := []string{}

    for i := range(old) {
        ret = append(ret, string(old[i]))
    }
    return ret
}

// string method cuz slices are funky
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    lettersPerDigit := map[byte]string {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    var (
        old, new []string
    )
    old = []string{""}

    for i := range(digits) {
        letters := lettersPerDigit[digits[i]]
        for j := range(letters) {
            for k := range(old) {
                new = append(new, old[k] + letters[j:j+1])
            }
        }
        old = new
        new = []string{}
    }

    return old
}

// with recursion O(n4^n)
// 4^n max possibilities because max of 4 letters per digit
// n work to build string each time
func letterCombinations(digits string) []string {
    lettersPerDigit := map[byte]string {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    combis := []string{}

    var backtrack func(string)

    backtrack = func(str string) {
        if len(str) == len(digits) {
            combis = append(combis, str)
            return
        }
        letters := lettersPerDigit[digits[len(str)]]
        for i := range letters {
            backtrack(str + letters[i:i+1])
        }
    }

    if len(digits) > 0 {
        backtrack("")
    }

    return combis
}

// with bytes which is kind of funky
func letterCombinations(digits string) []string {
    lettersPerDigit := map[byte]string {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    combis := []string{}
    combisBytes := [][]byte{}

    var backtrack func([]byte)

    backtrack = func(b []byte) {
        if len(b) == len(digits) {
            combisBytes = append(combisBytes, b)
            return
        }

        letters := lettersPerDigit[digits[len(b)]]
        for i := range letters {
            bCopy := make([]byte, len(b), len(b) + 1)
            copy(bCopy, b)
            bCopy = append(bCopy, letters[i])
            backtrack(bCopy)
        }
    }

    if len(digits) > 0 {
        backtrack([]byte{})

        for i := range combisBytes {
            combis = append(combis, string(combisBytes[i]))
        }
    }

    return combis
}

// using index of letters for each data so if you have digits = 273, then it goes from [0, 0, 0] to [3, 4, 3]
// O(n4^n)
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    lettersPerDigit := map[byte]string {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    combis := []string{}

    indices := make([]int, len(digits))
    boundaries := make([]int, len(digits))
    count := 1
    for i := range digits {
        numLetters := len(lettersPerDigit[digits[i]])
        boundaries[i] = numLetters - 1
        count *= numLetters
    }

    updateIndices := func(indices []int) []int {
        carry := 0
        idx := len(indices) - 1
        if indices[idx] + 1 > boundaries[idx] {
            carry = 1
            indices[idx] = 0
        } else {
            indices[idx]++
        }
        idx--

        for carry > 0 && idx >= 0 {
            if indices[idx] + 1 > boundaries[idx] && idx != 0 {
                carry = 1
                indices[idx] = 0
            } else {
                indices[idx] += 1
                carry = 0
            }
            idx--
        }

        return indices
    }
    
    for i := 0; i < count; i++ {
        bytes := make([]byte, len(indices))
        for j := range indices {
            bytes[j] = lettersPerDigit[digits[j]][indices[j]]
        }
        combis = append(combis, string(bytes))

        indices = updateIndices(indices)
    }

    return combis
}
