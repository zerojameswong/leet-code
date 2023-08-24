# tried bottom up, doesnt work using the 2D grid approach
class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        can_reach = [[0 for i in range(amount + 1)] for j in range(len(coins))]
        for i in range(len(coins)):
            can_reach[i][-1] = 1

        for i in range(len(coins) - 1, -1, -1):
            for j in range(amount - 1, -1, -1):
                coin_val = coins[i]
                reachable = False
                if j + coin_val <= amount:
                    reachable = j + coin_val == amount or can_reach[i][j + coin_val]
                if reachable:
                    can_reach[i][j] += 1

        return can_reach[0][0]

# only beats 28.5%
class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        memo = {}

        def find_ways(coin_idx, cumulative_val):
            coin_val = coins[coin_idx]

            if (coin_idx, cumulative_val) in memo:
                return memo[coin_idx, cumulative_val]

            ways = 0

            if coin_idx < len(coins) - 1:
                ways += find_ways(coin_idx + 1, cumulative_val)
            if cumulative_val + coin_val <= amount:
                ways += find_ways(coin_idx, cumulative_val + coin_val)
            if cumulative_val == amount:
                ways = 1

            memo[coin_idx, cumulative_val] = ways
            return ways

        return find_ways(0, 0)
        