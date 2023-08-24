# beats 95%
class Solution:
    def maximalNetworkRank(self, n: int, roads: List[List[int]]) -> int:
        max_network_rank = 0

        counts = [0 for i in range(n)]
        directly_connected = {}
        for road in roads:
            i, j = road
            if i > j:
                i, j = j, i
            counts[i] += 1
            counts[j] += 1
            directly_connected[i, j] = True

            if max_network_rank == 0:
                max_network_rank = 1

        for i in range(n):
            for j in range(i + 1, n):
                count = counts[i] + counts[j]
                if (i, j) in directly_connected:
                    count -= 1

                if count > max_network_rank:
                    max_network_rank = count

        return max_network_rank
