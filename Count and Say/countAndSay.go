// original beats 5.77% lol
func countAndSay(n int) string {
    var inner func(int, int, string) string
    inner = func(end, current int, s string) string {
        fmt.Println(end, current, s)
        if current == 0 {
            return inner(end, 1, "1")
        }

        if current == end {
            return s
        }

        count := 0
        counted := ""
        new := ""
        for i := range(s) {
            if s[i:i+1] != counted {
                if counted != "" {
                        new += strconv.Itoa(count)
                        new += counted
                }

                counted = s[i:i+1]
                count = 1
            } else {
                count++
            }
        }
        new += strconv.Itoa(count)
        new += counted
        return inner(end, current + 1, new)
    }

    return inner(n, 0, "")
}

// using a bytes buffer and fmt.Sprint i get to beats 69.62%
func countAndSay(n int) string {
    var inner func(int, int, string) string
    inner = func(end, current int, s string) string {
        fmt.Println(end, current, s)
        if current == 0 {
            return inner(end, 1, "1")
        }

        if current == end {
            return s
        }

        count := 0
        counted := ""
        new := bytes.Buffer{}
        for i := range(s) {
            if s[i:i+1] != counted {
                if counted != "" {
                    new.WriteString(fmt.Sprint(count))
                    new.WriteString(counted)
                }

                counted = s[i:i+1]
                count = 1
            } else {
                count++
            }
        }
        new.WriteString(fmt.Sprint(count))
        new.WriteString(counted)
        return inner(end, current + 1, new.String())
    }

    return inner(n, 0, "")
}
