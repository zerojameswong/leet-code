# bottom up approach, only beats 25% but the fastest entries are not DP
# they search for longer and longer palindromic substrings, both odd and even, 
# starting from all positions in the array 
class Solution:
    def longestPalindrome(self, s: str) -> str:
        longest_pal_start = 0
        longest_pal_end = 0
        longest_pal_len = 0

        is_pal = [[False for i in range(len(s))] for j in range(len(s))]

        for i in range(len(s)):
            is_pal[i][i] = True
            if longest_pal_len < 1:
                    longest_pal_start = i
                    longest_pal_end = i
                    longest_pal_len = 1

        for i in range(len(s) - 1):
            if s[i] == s[i + 1]:
                is_pal[i][i + 1] = True
                if longest_pal_len < 2:
                    longest_pal_start = i
                    longest_pal_end = i + 1
                    longest_pal_len = 2

        for length in range(3, len(s) + 1):
            for i in range(len(s) - length + 1):
                j = i + length - 1
                if is_pal[i + 1][j - 1] and s[i] == s[j]:
                    is_pal[i][j] = True
                    if longest_pal_len < length:
                        longest_pal_start = i
                        longest_pal_end = j
                        longest_pal_len = length

        return s[longest_pal_start : longest_pal_end + 1]

# memo top down, cant get it to work for long inputs
class Solution:
    def longestPalindrome(self, s: str) -> str:
        self.memo = {}
        self.longest_len = 0
        self.longest_l = 0
        self.longest_r = 0
        self.s = s

        self.lp(0, len(s) - 1)
        return s[self.longest_l : self.longest_r + 1]

    def lp(self, l, r):
        if r - l >= 2:
            if (l, r - 1) not in self.memo:
                self.lp(l, r - 1)
            if (l + 1, r) not in self.memo:
                self.lp(l + 1, r)

        eq = False
        length = r - l + 1

        if r == l:
            eq = True
        elif r - l <= 2:
            eq = self.s[l] == self.s[r]
            self.memo[l, r] = eq
        else:
            eq = self.memo[l + 1, r - 1] and self.s[l] == self.s[r]
            self.memo[l, r] = eq

        if eq and self.longest_len < length:
                self.longest_len = length
                self.longest_l = l
                self.longest_r = r
