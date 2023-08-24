// no brain backtrack version
func generateParenthesis(n int) []string {
    combis := []string{}

    var backtrack func([2]int, string)
    backtrack = func(budget [2]int, s string) {
        if len(s) == 2 * n {
            if isValid(s) {
                combis = append(combis, s)
            }
            return
        }
        if budget[0] > 0 {
            backtrack([2]int{budget[0] - 1, budget[1]}, s + "(")
        }
        if len(s) != 0 && budget[1] > 0 {
            backtrack([2]int{budget[0], budget[1] - 1}, s + ")")
        }
    }

    backtrack([2]int{n, n}, "")

    return combis
}

// i stole this from my previous leet code problem
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

// this is big brain version use the condition of no more budget means done 
// and only put a right bracket if there are more right brackets than left available
// to maintain the validity of the sequence so we dont need another function
func generateParenthesis(n int) []string {
    combis := []string{}

    var backtrack func([2]int, string)
    backtrack = func(budget [2]int, s string) {
        if budget[0] + budget[1] == 0 {
            combis = append(combis, s)
            return
        }
        if budget[0] > 0 {
            backtrack([2]int{budget[0] - 1, budget[1]}, s + "(")
        }
        if budget[1] > 0 && budget[1] > budget[0] {
            backtrack([2]int{budget[0], budget[1] - 1}, s + ")")
        }
    }

    backtrack([2]int{n, n}, "")

    return combis
}
