# beats 56%
class Solution:
    def repeatedSubstringPattern(self, s: str) -> bool:
        max_length = len(s) // 2
        for length in range(1, max_length + 1):
            if len(s) % length > 0:
                continue
            sub = s[:length]
            for i in range(length, len(s), length):
                if s[i:i+length] != sub:
                    break
            else:
                return True

        return False
