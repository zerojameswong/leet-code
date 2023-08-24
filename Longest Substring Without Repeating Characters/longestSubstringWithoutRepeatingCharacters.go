func lengthOfLongestSubstring(s string) int {
    var maxSubstringLength int = 1

    for start := 0; start < len(s) - 1; start++ {
        out:
        for end := start + 1; end < len(s); end++ {
            for compare := start; compare < end; compare++ {
                if s[compare] == s[end] {
                    substringLength := end - start
                    if substringLength > maxSubstringLength {
                        maxSubstringLength = substringLength
                    }
                    break out
                }
            }
        }    
    }

    return maxSubstringLength
}

// not really an improvement, uses O(n^2) space on top of that as well :(
func lengthOfLongestSubstring(s string) int {
    var maxSubstringLength int = 1
    differenceMatrix := make([][]bool, len(s))
    for i := range(s) {
        differenceMatrix[i] = make([]bool, len(s))
    }

    for i := 0; i < len(s); i++ {
        for j := 0; j < len(s); j++ {
            if s[i] == s[j] {
                differenceMatrix[i][j] = false
            } else {
                differenceMatrix[i][j] = true
            }
        }
    }

    for start := 0; start < len(s) - 1; start++ {
        out:
        for end := start + 1; end < len(s); end++ {
            different := true
            for compare := start; compare < end; compare++ {
                different = different && differenceMatrix[compare][end]
            }
            if !different {
                substringLength := end - start
                if substringLength > maxSubstringLength {
                    maxSubstringLength = substringLength
                }
                break out
            }
        }    
    }

    return maxSubstringLength
}

// O(n) with O(1) space complexity
func lengthOfLongestSubstring(s string) int {
    visited := make([]int, 256)

    longestSubstringLength := 0
    start := 0
    for idx, i := range(s) {
        if visited[i] >= start {
            start = idx
        }
        length :=  idx + 1 - start
        if length > longestSubstringLength {
            longestSubstringLength = length
        }

        visited[i] = idx
    }
    return longestSubstringLength
}
