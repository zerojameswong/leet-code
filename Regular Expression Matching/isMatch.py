class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        matched = False

        def backtrack(idx_string, idx_pattern):
            if idx_pattern < 0 and idx_string < 0:
                nonlocal matched
                matched = True
                return
            
            if idx_pattern < 0 or idx_string < 0 and p[idx_pattern] != "*":
                return
            
            if p[idx_pattern] == "*":
                backtrack(idx_string, idx_pattern - 2)

                to_match = p[idx_pattern - 1]
                if to_match == ".":
                    for i in range(idx_string, -1, -1):
                        backtrack(i - 1, idx_pattern - 2)
                else:
                    i = idx_string
                    while i >= 0 and s[i] == to_match:
                        backtrack(i - 1, idx_pattern - 2)
                        i -= 1
            elif p[idx_pattern] == "." or s[idx_string] == p[idx_pattern]:
                backtrack(idx_string - 1, idx_pattern - 1)
            elif s[idx_string] != p[idx_pattern]:
                return

        backtrack(len(s) - 1, len(p) - 1)
    
        return matched
