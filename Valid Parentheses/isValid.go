// my version, kinda clunky looking
func isValid(s string) bool {
    stack := []byte{}

    pairs := map[byte]byte {
        '(': ')',
        '[': ']',
        '{': '}',
    }

    for i := range s {
        if s[i] == '(' || s[i] == '[' || s[i] == '{' {
            stack = append(stack, pairs[s[i]])
        } else if len(stack) == 0 {
            return false
        } else {
            if s[i] == stack[len(stack) - 1] {
                stack = stack[:len(stack) - 1]
            } else {
                return false
            }
        }
    }
    if len(stack) == 0 {
        return true
    } else {
        return false
    }
}
