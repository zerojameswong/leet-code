// without using strconv, if you use strconv then this problem is trivial
func isPalindrome(x int) bool {
    digits := []int{}

    if x < 0 {
        return false
    }

    if x == 0 {
        return true
    }

    for x != 0 {
        digit := x % 10
        x /= 10
        digits = append(digits, digit)
    }

    left := 0
    right := len(digits) - 1

    isPalindrome := false
    for {
        if digits[left] != digits[right] {
            break
        }
        if right - left <= 1 {
            isPalindrome = true
            break
        }
        left++;
        right--;
    }

    return isPalindrome
}
