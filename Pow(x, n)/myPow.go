// cant do the braindead O(n) way cuz it takes too long, this way is O(logn)
func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    if x == 1 {
        return 1
    }

    var res float64

    if n > 0 {
        res = x
        i := 1
        for i * 2 <= n {
            res *= res
            i *= 2
        }
        remExponent := n - i

        for j := 0; j < remExponent; j++ {
            res *= x
        }
    }

    if n < 0 {
        res = 1 / x
        i := -1
        for i * 2 >= n {
            res *= res
            i *= 2
        }
        remExponent := -(n - i)

        for j := 0; j < remExponent; j++ {
            res *= 1 / x
        }
    }

    return res
}

