# bottom up beats 59%
class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        ways = [[0 for i in range(n)] for j in range(m)]
        ways[-1][-1] = 1

        for i in range(n - 1, -1, -1):
            for j in range(m - 1, -1, -1):
                if i == n - 1 and j == m - 1:
                    continue
                d, r = 0, 0
                if j + 1 < m:
                    d = ways[j + 1][i]
                if i + 1 < n:
                    r = ways[j][i + 1]
                ways[j][i] = d + r

        return ways[0][0]

# top down memoisation, beats 97%
class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        memo = {}
        memo[n - 1, m - 1] = 1

        def get_sum(x, y, n, m):
            if (x, y) in memo:
                return memo[x, y]
            
            d, r = 0, 0
            if x + 1 < n:
                d = get_sum(x + 1, y, n, m)
            if y + 1 < m:
                r = get_sum(x, y + 1, n, m)

            total = d + r
            memo[x, y] = total
            return total

        return get_sum(0, 0, n, m)
