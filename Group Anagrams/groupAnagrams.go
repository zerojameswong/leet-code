// O(nmlogm) where n is len(strs) and m is length of each individual string as sorting dominates the work in the loop and the number of loops is O(n)
// the other way to do this is to use map[[26]int][]string, the tradeoff is O(nm) time complexity but more memory usage
func groupAnagrams(strs []string) [][]string {
    anagramsPerKey := make(map[string][]string)

    for i := range strs {
        ints := make([]int, len(strs[i]))
        for j := range strs[i] {
            ints[j] = int(strs[i][j])
        }
        sort.Ints(ints)

        bytes := make([]byte, len(strs[i]))

        for j := range ints {
            bytes[j] = byte(ints[j])
        }

        k := string(bytes)
        anagramsPerKey[k] = append(anagramsPerKey[k], strs[i])
    }

    var anagrams [][]string
    for _, v := range anagramsPerKey {
        anagrams = append(anagrams, v)
    }

    return anagrams
}
