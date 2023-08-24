// find plane of reflection
// find palindromic extent
// this is my silly idea but it seems like O(n^2)
func longestPalindrome(s string) string {

    checks := [][2]int{}
    locations := make(map[int][]int)
    longestPalindromeLength := 1
    longestPalindrome := s[0:1]
    for idx, val := range(s) {
        locations[int(val)] = append(locations[int(val)], idx)
    }

    for _, v := range(locations) {
        if len(v) > 1 {
            for i := 0; i < len(v) - 1; i++ {
                if v[i + 1] - v[i] <= 2 {
                    checks = append(checks, [2]int{v[i], v[i + 1]})
                }
            }
        }
    }

    for _, check := range(checks) {
        length := 1
        done := false
        start := check[0]
        end := check[1]
        for !done {
            if start == 0 || end == len(s) - 1 {
                length = end - start + 1
                done = true
            } else if s[start - 1] != s[end + 1] {
                length = end - start + 1
                done = true
            } else {
                start -= 1
                end += 1
            }
        }
        if length > longestPalindromeLength {
            longestPalindromeLength = length
            longestPalindrome = s[start: end + 1]
        }
    }
    return longestPalindrome
}

// expand out from starting points, O(n^2)
// use indices so that we dont incur O(n) inner work from indexing the string which is O(n)
func longestPalindrome(s string) string {
    // odd length palindrome
    longestPalindromeLength := 0
    longestPalindromeIndex := [2]int{}
    for i := 0; i < len(s); i++ {
        start := i
        end := i
        for start >= 0 && end < len(s) && s[start] == s[end] {
            if end - start + 1 > longestPalindromeLength {
                longestPalindromeLength = end - start + 1
                longestPalindromeIndex = [2]int{start, end}
            }
            start -= 1
            end += 1
        }
    }

    // even length palindrome
        for i := 0; i < len(s) - 1; i++ {
        start := i
        end := i + 1
        for start >= 0 && end < len(s) && s[start] == s[end] {
            if end - start + 1 > longestPalindromeLength {
                longestPalindromeLength = end - start + 1
                longestPalindromeIndex = [2]int{start, end}
            }
            start -= 1
            end += 1
        }
    }

    longestPalindrome := s[longestPalindromeIndex[0] : longestPalindromeIndex[1] + 1]

    return longestPalindrome
}

// dynamic programming
// maintain a matrix where x[i][j] is true if the substring from index i to j is palindromic and false otherwise
// from length 1 to length of the whole string, take advantage of the pre-computed results for shorter substrings
func longestPalindrome(s string) string {
    startEndPalindromic := make([][]bool, len(s))
    for i := range(startEndPalindromic) {
        startEndPalindromic[i] = make([]bool, len(s))
    }

    longestPalindromeLength := 1
    longestPalindromeIndex := [2]int{0, 1}

    // length 1 palindromes
    for i := 0; i < len(s); i++ {
        startEndPalindromic[i][i] = true
    }

    // length 2 palindromes
    for i := 0; i < len(s) - 1; i++ {
        start := i
        end := i + 1
        if s[start] == s[end] {
            startEndPalindromic[start][end] = true
            if longestPalindromeLength < 2 {
                longestPalindromeLength = 2
                longestPalindromeIndex = [2]int{start, end}
            }
        }
    }

    for i := 3; i <= len(s); i++ {
        for j := 0; j <= len(s) - i; i++ {
            if startEndPalindromic[j + 1][j + i - 2] && s[j] == s[j + i - 1] {
                startEndPalindromic[j][j + i - 1] = true
                if i > longestPalindromeLength {
                    longestPalindromeLength = i
                    longestPalindromeIndex = [2]int{j, j + i - 1}
                }
            }
        }
    }

    longestPalindrome := s[longestPalindromeIndex[0] : longestPalindromeIndex[1] + 1]
    return longestPalindrome
}
