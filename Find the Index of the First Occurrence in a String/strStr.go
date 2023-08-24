func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) {
        return -1
    }
    if len(needle) == len(haystack) {
        if needle == haystack {
            return 0
        } else {
            return -1
        }
    }
    for i := 0; i < len(haystack) - len(needle) + 1; i++ {
        j := 0
        for ; j < len(needle); j++ {
            if haystack[i + j] != needle[j] {
                break
            }
        }
        if j == len(needle) {
            return i
        }
    }
    return -1
}