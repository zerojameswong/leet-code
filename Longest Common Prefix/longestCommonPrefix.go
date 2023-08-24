// my brutish approach
func longestCommonPrefix(strs []string) string {
    if len(strs) == 1 {
        return strs[0]
    }

    minLength := math.MaxInt32
    for i := range(strs) {
        if len(strs[i]) < minLength {
            minLength = len(strs[i])
        }
    }

    prefix := []byte{}
    out:
    for i := 0; i < minLength; i++ {
        reference := strs[0][i]
        for j := 1; j < len(strs); j++ {
            if strs[j][i] != reference {
                break out
            }
        }
        prefix = append(prefix, reference)
    }

    return string(prefix)
}
/*
 not my approach, takes advantage of the fact that the 0th string is either the shortest string or not
 whichever the shortest string is, its length limits the longest prefix
 if another string is the shortest substring then we can detect that too
 otherwise we keep going until there is a deviation
 by using a slice we get O(1) insertion, however if we did something like s := "" and s += char then we get an extra O(m^2) where m is the length of the shortest as we need to rebuild the string of length O(m) and this occurs in the outerloop which is executed o(m) times
 string(prefix) costs O(m) but this is not the dominant term which is O(mn)
*/
func longestCommonPrefix(strs []string) string {
    if len(strs) == 1 {
        return strs[0]
    }

    prefix := make([]byte, len(strs[0]))

    for i := 0; i < len(strs[0]); i++ {
        for j := 0; j < len(strs); j++ {
            if i == len(strs[j]) || strs[j][i] != strs[0][i] {
                return string(prefix)
            }
        }
        prefix = append(prefix, strs[0][i])
    }

    return string(prefix)
}
